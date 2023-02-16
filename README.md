# Overview

This project is a workaround example to create hybrid code to that support serverless or containerized way. This example code can serve multiple path in one function (normally only one endpoint per one function)

# Deployment

TBD

Deploy with runtime v2
```
gcloud functions deploy <YOUR_FUNCTION_NAME> \
--gen2 \
--runtime=go119 \
--project=<YOUR_PROJECT> \
--region=<YOUR_CHOOSEN_REGION> \
--source=. \
--entry-point=<YOUR_FUNCTION_NAME> \
--trigger-http \
--memory=128Mi \
--allow-unauthenticated 
```

# Test Case

TBD

# License

TBD