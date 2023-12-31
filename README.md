# sls_poc
[serverless framework](https://app.serverless.com) playground

### Installing serverless
```bash
npm install -g serverless
```

### Init project
```bash
serverless
```
Should result in following selection:
```bash
Creating a new serverless project

? What do you want to make? (Use arrow keys)
> AWS - Node.js - Starter
  AWS - Node.js - HTTP API
  AWS - Node.js - Scheduled Task
  AWS - Node.js - SQS Worker
  AWS - Node.js - Express API
  AWS - Node.js - Express API with DynamoDB
  AWS - Python - Starter
  AWS - Python - HTTP API
  AWS - Python - Scheduled Task
  AWS - Python - SQS Worker
  AWS - Python - Flask API
  AWS - Python - Flask API with DynamoDB
  Other
```

### List more templates
```bash
serverless create --help
```

### Create project from template example
```bash
serverless create --template aws-go --path golang_poc
```
```bash
serverless create --template aws-nodejs --path hd_wallet_js
```

### Invoke function
```bash
serverless invoke [local] --function <func_name>
```

### Invoke http triggered function using offline plugin
This simplifies development on local and doesn't need Docker
```bash
serverless offline
```

### Deploy function
```bash
serverless deploy --verbose  --function <func_name>
```

### Remove
```bash
serverless remove
```

## Build without deploy (updates .serveless dir)
```bash
serverless package
```