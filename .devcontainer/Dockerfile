FROM --platform=linux/amd64 pulumi/pulumi:3.123.0

RUN <<EOF
apt install -y bash-completion vim

## Set bash completion for pulumi
pulumi gen-completion bash >> /etc/bash_completion.d/pulumi

## Install LocalStack CLI
curl -Lo localstack-cli-2.3.0-linux-amd64-onefile.tar.gz \
    https://github.com/localstack/localstack-cli/releases/download/v2.3.0/localstack-cli-2.3.0-linux-amd64-onefile.tar.gz

tar xvzf localstack-cli-2.3.0-linux-*-onefile.tar.gz -C /usr/local/bin

## Install AWS CLI
pip install awscli

## Install Localstack AWS CLI
pip install awscli-local[ver1]
EOF
