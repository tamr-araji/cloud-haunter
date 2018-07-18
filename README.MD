[![CircleCI](https://circleci.com/gh/hortonworks/cloud-haunter.svg?style=shield)](https://circleci.com/gh/hortonworks/cloud-haunter) [![Go Report Card](https://goreportcard.com/badge/github.com/hortonworks/cloud-haunter)](https://goreportcard.com/report/github.com/hortonworks/cloud-haunter)

# Cloud Haunter

`Haunter is a pokémon ghost which will poison and kill your unidentified and orphan cloud resources`
 
The main goal of this project is to save costs in the cloud by monitoring and terminating resources.
Cloud Haunter (ch) is supports the three major cloud providers.

`The project is under construction, use on your own risk!`

Current version of the software has limited features, it can monitor and report (only HipChat supported)
 * long running instances
 * instances without owner
 * old cloud credentials

For the proper work, you have to use some custom tags/labels on your cloud resources.

#### AWS
 * Owner tag: `Owner`
 * Ignore tag: `cloud-cost-reducer-ignore`

#### Azure
 * Owner tag: `Owner`
 * Ignore tag: `cloud-cost-reducer-ignore`
 * Resource creation time, because Azure SDK doesn't response it: `cb-creation-timestamp`

#### Google
 * Owner label: `owner`
 * Ignore label: `cloud-cost-reducer-ignore`

There is an option to use custom tags/labels, but you have to recompile the project with a bunch of custom environment variables like `AWS_OWNER_LABEL`. Please check Makefile.

### Ignorance config

This tool works well if you use it from early days of your cloud account and all of your users are following the basic rules of tagging instances. On the other hand introducing it on an existing environment should be pain in the back.
There is an option to declare your ignore policies in a YAML file (please have look at utils/testdata/ignores.yml).
CH will skip all the resources where the name, owner, or any of the tags are starting with the given configuration.

## Future plans
 * automatic instance termination
 * more notification types
 * looking for orphan resources (images, volumes, IPs)
 * multiple cloud accounts

## Installation

`# curl -Ls https://github.com/hortonworks/cloud-haunter/releases/download/v0.2.1/cloud-haunter_0.2.1_$(uname)_x86_64.tgz | tar x`

## Usage

```
NAME:
   Cloud Haunter
USAGE:
   ch -o=operation -a=action [-c=cloud]
VERSION:
   0.2.0
AUTHOR(S):
   Hortonworks
OPERATIONS:
        -o longrunning
        -o oldaccess
        -o ownerless
ACTIONS:
        -a log
        -a notification
CLOUDS:
        -c AWS
        -c AZURE
        -c GCP
IGNORE:
        -i=/location/of/ignore/config.yml
DRY RUN:
        -d
VERBOSE:
        -v
HELP:
        -p
```

### Environment variables

#### AWS
 * AWS_ACCESS_KEY_ID
 * AWS_SECRET_ACCESS_KEY

#### Azure
 * AZURE_SUBSCRIPTION_ID
 * AZURE_TENANT_ID
 * AZURE_CLIENT_ID
 * AZURE_CLIENT_SECRET

#### Google
 * GOOGLE_PROJECT_ID
 * GOOGLE_APPLICATION_CREDENTIALS, location of service account JSON 

#### HipChat
 * HIPCHAT_TOKEN
 * HIPCHAT_SERVER
 * HIPCHAT_ROOM

#### Long running
 * RUNNING_PERIOD, default: 24h

#### Old access
 * ACCESS_AVAILABLE_PERIOD, default: 2880h