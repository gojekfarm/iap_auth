# IAP Auth  [![CircleCI](https://circleci.com/gh/gojekfarm/iap_auth.svg?style=svg)](https://circleci.com/gh/gojekfarm/iap_auth)

# Install as a Service

1. Add service account to IAP and download the json for service account credentials.
2. Create this kube secret
`kubectl create secret generic some-svc-sa-creds --from-file=sa.json="serviceaccountfiledownloadedfromgcp.json"`
3. Install as a service
`helm install gojektech-incubator/iap-auth --name=some-svc-iap --set iapHost=https://somehost,clientId=someclientid,secretName=some-svc-sa-creds`

# Install as a Sidecar


