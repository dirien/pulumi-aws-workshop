# Pulumi Workshop: Foundations to Advanced

[![Open in DevPod!](https://devpod.sh/assets/open-in-devpod.svg)](https://devpod.sh/open#https://github.com/dirien/pulumi-aws-workshop)

<img src="img/main.png">

## Welcome, Pulumi and Friends! 👋

The goal of this workshop is for you to learn how to use Pulumi in different scenarios. From running instances and
managed service to Kubernetes and GitOps. The workshop is split in multiple chapters showing you the different aspects
of Pulumi. As a bonus, you will learn how to use [AWS](https://aws.amazon.com)

### Repository

You can find the repository for this workshop [here](https://github.com/dirien/pulumi-aws-workshop). Please feel
free to look for the examples in the different chapters if you get stuck.

### Content

- [Chapter 0 - Hello, AWS World!](./00-hello-aws-world.md)
- [Chapter 1 - Welcome, EKS!](./01-eks-cluster-setup.md)
- [Chapter 2 - Containerize an Application](./02-app.md)
- [Chapter 3 - Deploy the Application to Kubernetes](./03-simple-deploy-app.md)
- [Chapter 4 - Enter the World of GitOps!](./04-argocd-setup.md)
- [Chapter 5 - Housekeeping!](./05-housekeeping.md)

### Prerequisites

You will need to install these tools in order to complete this workshop:

- [Pulumi](https://www.pulumi.com/docs/get-started/install/)
- [Pulumi Account](https://app.pulumi.com/signup) - this optional, but convenient to handle the state of the different
  stacks.
- [node.js](https://nodejs.org/en/download/)
- [Go](https://golang.org/doc/install)
- [Python](https://www.python.org/downloads/)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [A GitHub Account](https://github.com/signup)
- [Helm](https://helm.sh/docs/intro/install/)

There is also a [devcontainer.json](.devcontainer/devcontainer.json) file in this repository which you can use to spin
up a `devcontainer` with all the tools installed. Highly recommended if you are
using [VSCode](https://code.visualstudio.com/docs/devcontainers/containers), [GitHub Codespaces](https://docs.github.com/en/codespaces/overview)
or
[DevPods](https://devpod.sh).

### Install DevPod and the AWS DevPod provider

The best results you will get if you use [DevPods](https://devpod.sh) to run this workshop.

Select the Provider of your choice and configure it. You can find the documentation for the different
providers [here](https://devpod.sh/docs/managing-providers/add-provider).

Now you can add a new workspace by clicking on `Workspaces` -> `+ Create` and
enter `github.com/dirien/pulumi-aws-workshop` in the `Enter Workspace Source and click `Create Workspace`.

<img src="img/devpod_add_workspace.png">

### Troubleshooting Tips

If you encounter any challenges during the workshops, consider the following steps in order:

1. Don't hesitate to reach out to me! I'm always here to assist and get you back on track.
1. Review the example code available [here](https://github.com/dirien/pulumi-aws-workshop).
1. Search for the error on Google. Honestly, this method often provides the most insightful solutions.
1. Engage with the Pulumi Community on Slack. If you haven't joined yet, you can do
   so [here](https://slack.pulumi.com/).

### Want to know more?

If you enjoyed this workshop, please some of Pulumi's other [learning materials](https://www.pulumi.com/learn/)
