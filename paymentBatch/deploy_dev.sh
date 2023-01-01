#!/bin/bash

sam build --profile breathings -t ./template.yaml && sam deploy --profile breathings --config-env apne2-dev --no-progressbar --no-confirm-changeset
