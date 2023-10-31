package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")
		// Create VPC.
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsHostnames: pulumi.Bool(true),
			EnableDnsSupport:   pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create an internet gateway.
		gateway, err := ec2.NewInternetGateway(ctx, "gateway", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}

		// Create a subnet that automatically assigns new instances a public IP address.
		subnet, err := ec2.NewSubnet(ctx, "subnet", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.1.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create a route table.
		routeTable, err := ec2.NewRouteTable(ctx, "route-table", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: gateway.ID(),
				},
			},
		})
		if err != nil {
			return err
		}

		// Associate the route table with the public subnet.
		_, err = ec2.NewRouteTableAssociation(ctx, "route-table-association", &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet.ID(),
			RouteTableId: routeTable.ID(),
		})
		if err != nil {
			return err
		}

		ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			Filters: []ec2.GetAmiFilter{
				ec2.GetAmiFilter{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-lunar-23.04-amd64-server-*",
					},
				},
			},
			Owners: []string{
				"099720109477",
			},
			MostRecent: pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}

		// Create a security group allowing inbound access over port 8080 and outbound
		// access to anywhere.
		secGroup, err := ec2.NewSecurityGroup(ctx, "sec-group", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Enable HTTP access"),
			VpcId:       vpc.ID(),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					FromPort: pulumi.Int(cfg.GetInt("httpPort")),
					ToPort:   pulumi.Int(cfg.GetInt("httpPort")),
					Protocol: pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		server, err := ec2.NewInstance(ctx, "simple-vm", &ec2.InstanceArgs{
			InstanceType: pulumi.String("t3.micro"),
			SubnetId:     subnet.ID(),
			VpcSecurityGroupIds: pulumi.StringArray{
				secGroup.ID(),
			},
			UserData: pulumi.String(cfg.Require("userData")),
			Ami:      pulumi.String(ami.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("simple-vm"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("publicIp", server.PublicIp)
		return nil
	})
}
