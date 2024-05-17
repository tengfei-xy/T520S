#!/bin/bash

TOKEN_KEY_FILE_NAME=/Users/melta/Documents/source/script/testshell/token_key
# TOKEN_KEY_FILE_NAME=/home/tengfei/Documents/T520S/bark/token_key
# create_token_file=/home/tengfei/Documents/T520S/bark/authentication_token
create_token_file=/Users/melta/Documents/source/script/testshell/authentication_token
TEAM_ID=5U8LBRXG3A
AUTH_KEY_ID=LH4T9V5U4R

JWT_ISSUE_TIME=$(date +%s)
JWT_HEADER=$(printf '{ "alg": "ES256", "kid": "%s" }' "${AUTH_KEY_ID}" | openssl base64 -e -A | tr -- '+/' '-_' | tr -d =)
JWT_CLAIMS=$(printf '{ "iss": "%s", "iat": %d }' "${TEAM_ID}" "${JWT_ISSUE_TIME}" | openssl base64 -e -A | tr -- '+/' '-_' | tr -d =)
JWT_HEADER_CLAIMS="${JWT_HEADER}.${JWT_CLAIMS}"
JWT_SIGNED_HEADER_CLAIMS=$(printf "${JWT_HEADER_CLAIMS}" | openssl dgst -binary -sha256 -sign "${TOKEN_KEY_FILE_NAME}" | openssl base64 -e -A | tr -- '+/' '-_' | tr -d =)
AUTHENTICATION_TOKEN="${JWT_HEADER}.${JWT_CLAIMS}.${JWT_SIGNED_HEADER_CLAIMS}"

if [ -r "${create_token_file}" ];then
    os=$(uname)
    case $os in
    Darwin)
        s=$(stat -f %Sm "${create_token_file}")
        old_file_timestamp=$(date -j -f "%b %d %T %Y" "${s}" +%s)
        now_timestamp=$(date +%s)
        ;;
    Linux)
        s=$(stat -c %y "${create_token_file}")
        old_file_timestamp=$(date -d "${s}" +%s)
        now_timestamp=$(date +%s)
        ;;
    esac

    test $((now_timestamp-old_file_timestamp)) -le 3000 && { cat ${create_token_file}; exit 0; }
fi
echo "$AUTHENTICATION_TOKEN" | tee ${create_token_file}