#!/bin/bash
if [ -z "$AVVY_CLIENT_COMMON" ]; then
  client_common=./client-common
else
  client_common=$AVVY_CLIENT_COMMON
fi

cp $client_common/records/records.json avvy
