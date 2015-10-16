# terraform-provider-etcdiscovery

Terraform provider for generating
[etcd discovery tokens](https://coreos.com/os/docs/latest/cluster-discovery.html)

[![Circle CI](https://circleci.com/gh/paperg/terraform-provider-etcdiscovery.svg?style=svg)](https://circleci.com/gh/paperg/terraform-provider-etcdiscovery)

## Installation

  1. Install [Terraform](https://terraform.io/).
  2. `go get github.com/paperg/terraform-provider-etcdiscovery`

## Usage

    resource "etcdiscovery_token" "foo" { }
    resource "etcdiscovery_token" "bar" {
        size = 5
    }

    output "foo" {
        value = "${etcdiscovery_token.foo.id}"
    }

    output "bar" {
        value = "${etcdiscovery_token.bar.id}"
    }

Interpolate the `id` attribute into the cloud-init data for your CoreOS
instances for the win.

There isn't much else to it. Destroying the token doesn't really do anything
other that remove it from Terraform's state so that a new one will be generated
on the next run. This means you can `terraform apply` to create a CoreOS
cluster, then `terraform destroy` to kill it all, then `terraform apply` to
make a new one again and the new cluster will get a new discovery token, which
is probably what you wanted.

If you are running your own discovery service then you can configure an
alternate URL like this:

    provider "etcdiscovery" {
        url = "http://example.com/new"
    }
