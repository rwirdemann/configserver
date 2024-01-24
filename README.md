# Simple Config Server
Some endpoints to maintain simple key / value config items.

## Deployment
```
make deploy
```

## Endpoints

### GET /config
Returns config value for the given key as sting in the response body.

```
curl https://jkyszaoly5.execute-api.eu-central-1.amazonaws.com/dev/config/{key}
```

### POST /config
Creates or updates an existing config key value.

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