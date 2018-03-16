cmattoon/aws-env
================

Exports AWS meta-data as environment vars

    docker pull cmattoon/aws-env:latest


Example Output
```
AMI_ID="ami-xxxxxx"
AMI_LAUNCH_INDEX="0"
AMI_MANIFEST_PATH="(unknown)"
HOSTNAME="ip-172-1-1-1.us-west-2.compute.internal"
INSTANCE_ACTION="none"
INSTANCE_ID="i-xxxxxxxxxxxxxxx"
INSTANCE_TYPE="t2.medium"
LOCAL_HOSTNAME="ip-172-1-1-1.us-west-2.compute.internal"
LOCAL_IPV4="172.1.1.1"
MAC="aa:bb:cc:dd:ee:ff"
PLACEMENT/AVAILABILITY_ZONE="us-west-2a"
PROFILE="default-hvm"
PUBLIC_HOSTNAME="ec2-2-2-2-2.us-west-2.compute.amazonaws.com"
PUBLIC_IPV4="2.2.2.2"
RESERVATION_ID="r-xxxxxxxxxxxxx"
SECURITY_GROUPS="TestSecurityGroup"
```


### Use Case: Eval to set Environment

    eval $(docker run --rm -t cmattoon/aws-env:latest)


### Use Case: Write Environment File

    docker run --rm -e AWS_ENVIRONMENT_FILE=/foo/aws.env -v /tmp:/foo -t cmattoon/aws-env:latest


## Build

    make container

