package eks

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EKSCluster struct {
	pulumi.ResourceState

	Kubeconfig pulumi.StringOutput `pulumi:"kubeconfig"`
}

type EKSClusterArgs struct {
	ClusterName       string
	KubernetesVersion string
	InstanceType      string
	DesiredCapacity   int
	MinSize           int
	MaxSize           int
}

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

func NewEKSCluster(ctx *pulumi.Context, name string, args *EKSClusterArgs, opts ...pulumi.ResourceOption) (*EKSCluster, error) {
	var cluster EKSCluster
	err := ctx.RegisterComponentResource("pkg:eks:EKSCluster", name, &cluster, opts...)
	if err != nil {
		return nil, err
	}
	vpc, err := ec2.NewVpc(ctx, fmt.Sprintf("%s-vpc", name), &ec2.VpcArgs{
		CidrBlock:          pulumi.String("172.31.0.0/16"),
		EnableDnsHostnames: pulumi.Bool(true),
		EnableDnsSupport:   pulumi.Bool(true),
	}, pulumi.Parent(&cluster))
	if err != nil {
		return nil, err
	}

	// Create an internet gateway.
	gateway, err := ec2.NewInternetGateway(ctx, fmt.Sprintf("%s-gateway", name), &ec2.InternetGatewayArgs{
		VpcId: vpc.ID(),
	}, pulumi.Parent(&cluster))
	if err != nil {
		return nil, err
	}

	// Create a route table.
	routeTable, err := ec2.NewRouteTable(ctx, fmt.Sprintf("%s-route-table", name), &ec2.RouteTableArgs{
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
		publicSubnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("%s-subnet-%d", name, i), &ec2.SubnetArgs{
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
		_, err = ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("%s-route-table-association-%d", name, i), &ec2.RouteTableAssociationArgs{
			RouteTableId: routeTable.ID(),
			SubnetId:     publicSubnet.ID(),
		})
		if err != nil {
			return nil, err
		}
		publicSubnetIDs = append(publicSubnetIDs, publicSubnet.ID())
	}

	eksCluster, err := eks.NewCluster(ctx, fmt.Sprintf("%s-eks-cluster", name), &eks.ClusterArgs{
		Name:            pulumi.String(args.ClusterName),
		VpcId:           vpc.ID(),
		SubnetIds:       publicSubnetIDs,
		Version:         pulumi.String(args.KubernetesVersion),
		InstanceType:    pulumi.String(args.InstanceType),
		DesiredCapacity: pulumi.Int(args.DesiredCapacity),
		MinSize:         pulumi.Int(args.MinSize),
		MaxSize:         pulumi.Int(args.MaxSize),
		ProviderCredentialOpts: eks.KubeconfigOptionsArgs{
			ProfileName: pulumi.String("default"),
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("eks"),
		},
	}, pulumi.Parent(&cluster))
	if err != nil {
		return nil, err
	}

	cluster.Kubeconfig = eksCluster.KubeconfigJson

	if err := ctx.RegisterResourceOutputs(&cluster, pulumi.Map{
		"kubeconfig": cluster.Kubeconfig,
	}); err != nil {
		return nil, err
	}

	return &cluster, nil
}
