# Hugo Amplify Test

This repo contains a static frontend based on Hugo which can talk to a backend
to support filling out a contact form; the backend is realized as a Lambda function
which sends an email via AWS SES.

This was essentially a toy project to look at how to glue these things together;
in practice, there are significant constraints around using SES and having a
HTTP endpoint hook into SES email delivery should be managed with care.

That said, if you want to set it up, instructions below.

As you will need to register a webhook with the repo, you will need to make a
fork of this to deploy it to Amplify.

## Setting up the backend

Prerequisites:
- go build tools (I used v1.14)
- SAM tooling
- Configured AWS account

Building and deploying the backend is straightforward; to build the lambda function,
use the build script provided.

```
cd backend
./build.sh
```

To deploy the lambda function, use the sam tooling:

```
sam deploy --guided -t backend-deployment.yaml
```

Once this terminates successfully, you have a working backend. You can test
as follows:

```
ADD CURL
```

## Setting up the frontend

Although AWS Amplify offers a CLI, here the web interface was used to manage
the application in Amplify.

### Connect the repo to Amplify

Connecting the repo to Amplify is straightforward; this registers a webhook with
the repo
In this case, we used the AWS
Setting up the frontend is more complex with multiple moving parts.

Firs
project which serves a simple site and provides a
form backend implemented as a Lambda function. It is referenced in this post
which describes how to use AWS Amplify to deploy such a Hugo website.

