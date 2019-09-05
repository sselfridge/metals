#!/bin/bash

GOOS=linux go build -o main main.go 
zip deployment.zip main
aws lambda update-function-code --region us-west-2 --function-name MeinGo --zip-file fileb://./deployment.zip
