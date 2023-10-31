# Chapter 1 - Welcome, EKS!

<img src="img/chap2.png">

## Overview

In this chapter, we'll set up a Scalable Kubernetes Service (SKS) cluster, which will serve as the foundation for our
subsequent workshop chapters.

While I'll be using `typescript` for this chapter, please choose the language you're most at ease with.

## Instructions

### Step 1 - Configure the AWS CLI

We're going to use the environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` to configure the AWS CLI.
Simply run the following commands to set them.

```bash
export AWS_ACCESS_KEY_ID=<your-access-key-id>
export AWS_SECRET_ACCESS_KEY=<your-secret-access-key>
```

To verify that the configuration is correct, run the following command.

```bash
aws sts get-caller-identity
```

### Step 2 - Configure the Pulumi CLI

> If you run Pulumi for the first time, you will be asked to log in. Follow the instructions on the screen to
> login. You may need to create an account first, don't worry it is free.

To initialize a new Pulumi project, run `pulumi new` and select from all the available templates the `typescript`. Of
course, you can use any other language you want.

```bash
pulumi new aws-go
```

You will be guided through a wizard to create a new Pulumi project. You can use the following values:

```bash
project name (01-eks-cluster-setup): 01-eks-cluster-setup
project description (A minimal AWS Go Pulumi program):  
Created project '01-eks-cluster-setup'

Please enter your desired stack name.
To create a stack in an organization, use the format <org-name>/<stack-name> (e.g. `acmecorp/dev`).
stack name (dev): dev 
...
aws:region: The AWS region to deploy into (us-east-1): eu-central-1 
```

The template `aws-go` will create a new Pulumi project with
the [Pulumi AWS provider](https://www.pulumi.com/registry/packages/aws/) already installed. For detailed instructions,
refer to the Pulumi AWS Provider documentation.

Use `config` values right from the outset for the following properties:

- `kubernetesVersion`
- `instanceType`
- `desiredCapacity`
- `minSize`
- `maxSize`
- `region`

Remember the techniques we discussed in the previous chapter!

To keep the creation very simple, we will use the `pulumi-eks` component resource. To add it to your project, run the
following command:

```bash
go get github.com/pulumi/pulumi-eks/sdk@v1.0.4
```

See the [Pulumi EKS Component Resource](https://www.pulumi.com/registry/packages/eks/) documentation for more
information.

Ensure you're using the latest version of Kubernetes. To determine the exact version, execute the
command `aws eks describe-addon-versions | jq -r ".addons[] | .addonVersions[] | .compatibilities[] | .clusterVersion" | sort | uniq`
versions and select the most recent version.

To wrap things up, export the `kubeconfig` from the `Cluster` resource. We'll require this later to establish a
connection to the cluster.

### Step 3 - Configure Kubectl

With the `pulumi stack output` command, you can retrieve any output value from the stack. In this case, we are going to
retrieve the kubeconfig to use with `kubectl`.

```bash
pulumi stack output kubeconfig --show-secrets -s dev > kubeconfig
```

### Step 4 - Verify the cluster

Now that we have the kubeconfig, we can verify the cluster is up and running. Not that we need this, but it is always
good to verify.

```bash
kubectl --kubeconfig kubeconfig get nodes
```

You should see a similar output:

```bash
NAME                                           STATUS   ROLES    AGE   VERSION
ip-172-31-15-2.eu-central-1.compute.internal   Ready    <none>   34s   v1.28.1-eks-43840fb
```

### Step 5 - Understanding the Need for Component Resources!

First, let's clean up by destroying the current stack:

```
pulumi destroy -s dev
```

Then, delete the stack using:

```
pulumi stack rm dev
```

We'll recreate the stack, but this time we'll utilize a component resource.

If you've observed, setting up an SKS cluster with all its requisite resources can be a tad intricate and isn't immune
to errors. The process demands a deep understanding, and every detail matters.

Component resources offer a solution. They serve as logical groupings of resources. Typically, components instantiate a
related set of resources in their constructor, treating them as children. This abstraction conceals the nitty-gritty,
ensuring a smoother experience.

For a deep dive, refer to the [Pulumi Component Resource](https://www.pulumi.com/docs/concepts/resources/components/)
documentation.

I've prepared a reference implementation for you in `internal/eks/eks.go`. Use it as a guide or starting point.

With the new component resource in place, you're all set to deploy the stack.

```bash
pulumi up
```

Congratulations! You have successfully deployed a Kubernetes cluster on AWS using Pulumi. Please leave the cluster
up and running for [Chapter 1 - Containerize an Application](./01-app-setup.md)

## Stretch Goals

- Are you looking to diversify your cluster? Consider adding a second node group with a distinct node type to the
  existing cluster.

## Learn More

- [Pulumi](https://www.pulumi.com/)
- [AWS EKS](https://aws.amazon.com/eks/)
- [Pulumi Component Resources](https://www.pulumi.com/docs/concepts/resources/components/)
- [Pulumi AWS provider](https://www.pulumi.com/registry/packages/aws/)
