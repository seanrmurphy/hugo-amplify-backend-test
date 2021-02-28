# Hugo Amplify Test

This repo contains a static frontend based on Hugo which can talk to a backend
to support filling out a contact form; the backend is realized as a Lambda
function which sends an email via AWS SES.

This was essentially a toy project to look at how to glue these things together;
it was noted that SES does have significant constraints, particularly around sending
emails to non-verified addresses when operating in sandbox mode - operating in non
sandbox mode requires AWS approval. These constraints, however, were not a blocker
here as all that was necessary was to be able to send an email to a single verified
email address when someone fills out the contact form.

For more information, there is a medium post on this project [here](https://seanrmurphy.medium.com/hugo-aws-amplify-and-a-sneaky-lambda-backend-4b5c47ba8382).

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

This will give you a backend which can send an email via SES. Due to the SES
sandbox constraints, however, it is necessary to ensure that a verified email
address is added to your SES sandbox; AWS provides instructions
[here](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses-procedure.html). Modify the
backend code to use this email address as it is currently hardcoded into the
application.

With this done, you should have a working backend. You can test as follows:

```
$ curl -X POST https://<endpoint>.execute-api.eu-west-1.amazonaws.com/contact -d 'email=tom@test.com&name=Tom&message=hello&_next=return'
{"message": "Will respond to email addr asap..."}
```

You should then receive an email to the verified email address.

## Setting up the frontend

Although AWS Amplify offers a CLI, here the web interface was used to manage
the application in Amplify.

### Connect the repo to Amplify

Connecting the repo to Amplify is straightforward; this registers a webhook with
the repo so you will need to have your own fork of this repo.

As the frontend content is in the `frontend` directory and the build process
requires go modules, you will need to modify the Build Settings slightly -
the `extras/amplify.yml` shows what should be in the Build Settings. With this,
you can use the standard Amazon Linux 2 image for building the hugo site.

You will need to make two specific modifications to the site, one to tell the
frontend that it is being served from Amplify and one to tell it where the
contact form backend is. For the first, it is necessary to modify the
`baseURL` parameter in `frontend/config/_default/config.toml` to point at the
endpoint provided by Amplify (or an FQDN if you connect that to this site). For
the second, it is necessary to modify the `amplify_id` parameter in the
`data/tnd-forms/contact.yaml` file to point to the backend.

When you commit and push these changes, Amplify should receive a webhook
notification and deploy a frontend with a contact form which triggers the
Lambda function to send an email to the specified email address.

