#!/bin/bash

# Please Change the following Values

if [[ -z "${DOCAREA_DOCUMENTATION_ID}" ]]; then
  DOCUMENTATION_ID=""
else
  DOCUMENTATION_ID="${DOCAREA_DOCUMENTATION_ID}"
fi


if [[ -z "${DOCAREA_CLIENT_ID}" ]]; then
  CLIENT_ID=""
else
  CLIENT_ID="${DOCAREA_CLIENT_ID}"
fi


if [[ -z "${DOCAREA_CLIENT_SECRET}" ]]; then
  CLIENT_SECRET=""
else
  CLIENT_SECRET="${DOCAREA_CLIENT_SECRET}"
fi

API_ENDPOINT="https://www.docarea.io"

# DO NOT CHANGE ANYTHING BELOW

uploaddoc=$1

echo "Calculate Size"
SIZE=$(du -sb ${uploaddoc} | awk '{print $1}')
echo "Build Meta Dependencies"

ARCHIVE_NAME=$(timeout --foreground 5s cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

if [[ -z "${ARCHIVE_NAME}" ]]; then
  ARCHIVE_NAME="CANTACCESSURANDOM"
fi

TEMP_UPLOAD_DIR=$(mktemp -d -t docarea-XXXXXXXXXX)
echo "Compress Documentation"
cd $uploaddoc; tar --xz -cf ${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea --exclude ".." * ./.*
echo "Build Checksum"
ARCHIVE_CHECKSUM=$(sha256sum ${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea | awk '{print $1}')

#  --data '{ "grant_type": "client_credentials", "scope": "read", "CLIENT_ID": "'$CLIENT_ID'", "CLIENT_SECRET": "'$CLIENT_SECRET'" }' \
#--header "Content-Type: application/json"
echo "Request permission"
tokenrequest=$(curl -s \
  -X POST \
  -d 'grant_type=client_credentials&scope=upload_documentation&client_id='$CLIENT_ID'&client_secret='$CLIENT_SECRET \
  ${API_ENDPOINT}/oauth2/token/)

token=$(echo $tokenrequest | jq -r '.access_token')

echo "Announce upload"
uploadtokenrequest=$(curl -s -H 'Content-Type: application/json' \
 -H "Authorization: Bearer ${token}" \
 -X POST \
 --data '{"state": "ok", "code": 200, "object": { "documentationId": "'${DOCUMENTATION_ID}'", "size": '${SIZE}', "checksum":"'${ARCHIVE_CHECKSUM}'", "sendMeta": false}}' \
 ${API_ENDPOINT}/api/upload/request)
 
echo $uploadtokenrequest

uploadtoken=$(echo "${uploadtokenrequest}" | jq -r '.object.uploadToken')

echo "Upload Archive"
curl  -s --request POST \
-F "documentation=@${TEMP_UPLOAD_DIR}/${ARCHIVE_NAME}.docarea" \
-H "Authorization: Bearer ${token}" \
"${API_ENDPOINT}/api/upload/${uploadtoken}"

echo "Done"
