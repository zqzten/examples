# conversion-webhook

This is minimal Kubernetes CRD conversion webhook sample.

## What's in it

It will install a CRD called `NetworkAddress` with two versions `v1` and `v2` into your cluster.

It will also install a conversion webhook for the CRD, the conversion logic can be found in `api/v2/networkaddress_conversion.go`.

No controller logic will be installed.

## How to install

1. Make sure your Kubernetes cluster has [cert-manager](https://cert-manager.io/) v1.x installed (for webhook cert auto signing)
2. Run `make install deploy`

## Other notes

If you don't trust or cannot pull my pre-built image in `Makefile`, you can run `make docker-build` to build your own from source and replace the `IMG` in `Makefile`.
