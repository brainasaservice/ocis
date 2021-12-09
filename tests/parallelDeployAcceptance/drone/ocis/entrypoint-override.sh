#!/bin/sh
set -e
ocis server &
sleep 10

# It is nice to have the following services stopped
# but currently, follwing commands are failing
# idp, glauth and accounts are not needed -> replaced by Keycloak and OpenLDAP

# ocis kill idp
# ocis kill accounts
# ocis kill glauth

wait
