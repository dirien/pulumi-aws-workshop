package main

import (
	"01-eks-cluster-setup/internal/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

/*
var (
	publicSubnetCidrs = []string{
		"172.31.0.0/20",
		"172.31.48.0/20",
	}
	availabilityZones = []string{
		"eu-central-1a",
		"eu-central-1b",
	}
)
*/

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		cfg := config.New(ctx, "")
		eksCluster, err := eks.NewEKSCluster(ctx, "eks", &eks.EKSClusterArgs{
			ClusterName:       "eks",
			KubernetesVersion: cfg.Get("kubernetesVersion"),
			InstanceType:      cfg.Get("instanceType"),
			DesiredCapacity:   cfg.GetInt("desiredCapacity"),
			MinSize:           cfg.GetInt("minSize"),
			MaxSize:           cfg.GetInt("maxSize"),
		})
		if err != nil {
			return err
		}
		ctx.Export("kubeconfig", pulumi.ToSecret(eksCluster.Kubeconfig))
		// Create VPC.
		/*vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("172.31.0.0/16"),
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

		// Create a route table.
		routeTable, err := ec2.NewRouteTable(ctx, "route-table", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: gateway.ID(),
				},
			},
		}, pulumi.Parent(&cluster))
		if err != nil {
			return nil, err
		}

		var publicSubnetIDs pulumi.StringArray

		// Create a subnet for each availability zone
		for i, az := range availabilityZones {
			publicSubnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("public-subnet-%d", i), &ec2.SubnetArgs{
				VpcId:                       vpc.ID(),
				CidrBlock:                   pulumi.String(publicSubnetCidrs[i]),
				MapPublicIpOnLaunch:         pulumi.Bool(true),
				AssignIpv6AddressOnCreation: pulumi.Bool(false),
				AvailabilityZone:            pulumi.String(az),
				Tags: pulumi.StringMap{
					"Name": pulumi.Sprintf("eks-public-subnet-%d", az),
				},
			})
			if err != nil {
				return nil, err
			}
			_, err = ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("route-table-association-%d", i), &ec2.RouteTableAssociationArgs{
				RouteTableId: routeTable.ID(),
				SubnetId:     publicSubnet.ID(),
			})
			if err != nil {
				return nil, err
			}
			publicSubnetIDs = append(publicSubnetIDs, publicSubnet.ID())
		}

		cluster, err := eks.NewCluster(ctx, "eks", &eks.ClusterArgs{
			Name:            pulumi.String("eks"),
			VpcId:           vpc.ID(),
			SubnetIds:       publicSubnetIDs,
			Version:         pulumi.String(cfg.Get("kubernetesVersion")),
			InstanceType:    pulumi.String(cfg.Get("instanceType")),
			DesiredCapacity: pulumi.Int(cfg.GetInt("desiredCapacity")),
			MinSize:         pulumi.Int(cfg.GetInt("minSize")),
			MaxSize:         pulumi.Int(cfg.GetInt("maxSize")),
			ProviderCredentialOpts: eks.KubeconfigOptionsArgs{
				ProfileName: pulumi.String("default"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("eks"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("kubeconfig", pulumi.ToSecret(cluster.KubeconfigJson))
		*/
		return nil
	})
}
