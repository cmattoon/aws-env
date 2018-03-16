cmattoon/aws-env
================

Exports AWS meta-data as environment vars

    docker pull cmattoon/aws-env:latest


### Use Case: Eval to set Environment

    eval $(docker run --rm -t cmattoon/aws-env:latest)


### Use Case: Write Environment File

    docker run --rm -e AWS_ENVIRONMENT_FILE=/foo/aws.env -v /tmp:/foo -t cmattoon/aws-env:latest


## Build

    make container

