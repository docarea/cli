#!/bin/bash

# Please Change the following Values

DOCUMENTATION_ID=""

CLIENT_ID=""
CLIENT_SECRET=""
API_ENDPOINT="http://www.docarea.io"

# DO NOT CHANGE ANYTHING BELOW

uploaddoc=$1


SIZE=$(du -sb ${uploaddoc} | awk '{print $1}')
ARCHIVE_NAME=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
TEMP_UPLOAD_DIR=$(mktemp -d -t docarea-XXXXXXXXXX)

cd $uploaddoc; tar --xz -cf ${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea *

ARCHIVE_CHECKSUM=$(sha256sum ${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea | awk '{print $1}')

#  --data '{ "grant_type": "client_credentials", "scope": "read", "CLIENT_ID": "'$CLIENT_ID'", "CLIENT_SECRET": "'$CLIENT_SECRET'" }' \
#--header "Content-Type: application/json"

tokenrequest=$(curl -s \
  -X POST \
  -d 'grant_type=client_credentials&scope=upload_documentation&client_id='$CLIENT_ID'&client_secret='$CLIENT_SECRET \
  ${API_ENDPOINT}/oauth2/token/)

token=$(echo $tokenrequest | jq -r '.access_token')


uploadtokenrequest=$(curl -s -H 'Content-Type: application/json' \
 -H "Authorization: Bearer ${token}" \
 -X POST \
 --data '{"state": "ok", "code": 200, "object": { "documentationId": "'${DOCUMENTATION_ID}'", "size": '${SIZE}', "checksum":"'${ARCHIVE_CHECKSUM}'", "sendMeta": false}}' \
 ${API_ENDPOINT}/api/upload/request)

uploadtoken=$(echo "${uploadtokenrequest}" | jq -r '.object.uploadToken')

echo "Upload Archive"
curl  -s --request POST \
--data-binary "@${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea" \
-H "Authorization: Bearer ${token}" \
-H 'Content-Type: application/vnd.docarea+archive' \
"${API_ENDPOINT}/api/upload/${uploadtoken}"

echo ${TEMP_UPLOAD_DIR}
# rm ${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea
