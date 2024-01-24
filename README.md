# Simple Config Server

## Deployment

```
make deploy
```

## Endpoints

### GET /config
```
curl https://jkyszaoly5.execute-api.eu-central-1.amazonaws.com/dev/config

# HTTP Status Codes
200: config key exsits
```

### POST /config
```
curl -X POST https://xo94oqzj8e.execute-api.eu-central-1.amazonaws.com/dev/config \
   -H 'Content-Type: application/json' \
   -d '{
      "key": "jobdog.publishservice.url",
      "value": "https://77srys74sh.execute-api.eu-central-1.amazonaws.com/dev/jobs"
   }'
```

## AWS CLI Hints
```
aws lambda list-functions
aws lambda get-function-configuration --function-name config-dev-putconfig
```

```
saw watch /aws/lambda/config-dev-putconfig
```

```
aws dynamodb scan --table-name ConfigItems
```