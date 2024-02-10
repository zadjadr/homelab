# Infrastructure

Use `pass edit cloud.oracle.com/auth` to create a secret with the following keys & fill them with the correct values:

```ini
TF_VAR_fingerprint=
TF_VAR_private_key_path=
TF_VAR_region=
TF_VAR_tenancy_ocid=
TF_VAR_user_ocid=
TF_VAR_compartment_ocid=
TF_VAR_ssh_authorized_keys=
TF_VAR_subnet_id=
TF_VAR_ipv6_cidr=
```

```shell
# Optional
export OCI_AZONES="ldMg:EU-FRANKFURT-1-AD-1 ldMg:EU-FRANKFURT-1-AD-2 ldMg:EU-FRANKFURT-1-AD-3"
export OCI_INSTANCE_NAMES="1.cloud"

cd infrastructure
go build -o checker

./checker
```
