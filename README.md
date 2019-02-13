# IAP Auth  [![CircleCI](https://circleci.com/gh/gojekfarm/iap_auth.svg?style=svg)](https://circleci.com/gh/gojekfarm/iap_auth)

## IAP Enabled Google Load Balancer
IAP: [Identity Aware Proxy] (https://cloud.google.com/iap/)
Read more about IAP here https://cloud.google.com/blog/products/identity-security/protecting-your-cloud-vms-with-cloud-iap-context-aware-access-controls

### TLDR;
1. Setup an https Google load balancer
2. Enable IAP (Security > Identity Aware Proxy)
   All eligible proxies will be listed here. IAP toggle will enable Oauth Bearer token based auth.
3. After enabling and selecting this you can add previously created service accounts to this proxy.
4. Download this service account credentials and configure as a param in in install

## Install as a Service

1. Add service account to IAP and download the json for service account credentials.
2. Create this kube secret

`kubectl create secret generic some-svc-sa-creds --from-file=sa.json="serviceaccountfiledownloadedfromgcp.json"`

3. Install as a service

`helm install gojektech-incubator/iap-auth --name=some-svc-iap --set iapHost=https://somehost,clientId=someclientid,secretName=some-svc-sa-creds`

## Dev setup

For go1.11, you need an environment variable set to enable [go modules](https://github.com/golang/go/wiki/Modules)

```
$ export GO111MODULE=on
```

Assuming you are in the directory `iap_auth`

### Running the tests

```
$ make setup
$ make test
```

### Building the binary

```
$ make build
# the compiled binary would be inside iap_auth/out/

```

### Running the binary

```
$ make copy-config
$ make setup
$ ./out/iap_auth server
```

## Install as a Sidecar

TODO
