#!/bin/sh

fs=""
if [[ "$1" == "s3" ]]
then
  echo "Using aws s3"
  # export AWS_LAMBDA_FUNCTION_NAME=$1
  profile=ian
  region=$(eval "aws configure get region --profile ${profile}")
  appid=$(eval "aws configure get aws_access_key_id --profile ${profile}")
  secret=$(eval "aws configure get aws_secret_access_key --profile ${profile}")
  #accountId=$(eval "aws configure get account_id --profile ${profile}")
  echo "region=$region"
  export aws_region=$region
  export aws_secret=$secret
  export aws_appid=$appid
  export ian_pkv=$(jq -r .privateKey test/923cef14-e6af-4c2d-999e-1396010b2b86.pri.ian)
  export aws_bucket="test-1-ian"
  fs="{\"fsType\": \"s3\"}"  
fi

cmd="./ian-test-0 --port=9312 --pkf=./test/923cef14-e6af-4c2d-999e-1396010b2b86.pri.ian --fs='$fs'"
echo "Running $cmd"
eval "$cmd"

