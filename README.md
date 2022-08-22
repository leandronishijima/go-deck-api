# Go Deck API

## API Reference

### Create a new deck

```http
  POST /api/deck/new
```

| Name     | Required | Description                         | Default value | Example             |
| -------- | -------- | ----------------------------------- | ------------- | ------------------- |
| shuffled | optional | The deck to be shuffled or not      | false         | true                |
| cards    | optional | Cards that will be used in the deck | -             | ["AS", "QS", "10H"] |

- Example request:
```curl
curl -X "POST" "http://localhost:8080/api/deck/new" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "cards": [
    "AS",
    "QS",
    "10H"
  ],
  "shuffled": true
}
```

- Response:
```json
{
  "deck_id": "3afb3e6f-b686-4cb8-8211-c6f58270453e",
  "remaining": 52,
  "shuffled": false
}
```

### Open a deck

```http
  GET /api/deck/open/${deck_id}
```

- Example request
```curl
curl "http://localhost:8080/api/deck/open/deck_id_uuid"
```

- Response:
```json
{
  "deck_id": "2d4ee730-94cc-4585-b02d-47238aee59f9",
  "shuffled": true,
  "remaining": 3,
  "cards": [
    {
      "value": "ACE",
      "suit": "SPADES",
      "code": "AS"
    },
    {
      "value": "10",
      "suit": "HEARTS",
      "code": "10H"
    },
    {
      "value": "QUEEN",
      "suit": "SPADES",
      "code": "QS"
    }
  ]
}
```

### Draw cards

```http
  PATCH /api/deck/${deck_id}/draw
```

- Example request
```curl
curl -X "PATCH" "http://localhost:8080/api/deck/deck_id_uuid/draw" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "count": 1
}'
```

- Response:
```json
{
  "cards": [
    {
      "value": "3",
      "suit": "SPADES",
      "code": "3S"
    },
    {
      "value": "4",
      "suit": "SPADES",
      "code": "4S"
    },
    {
      "value": "5",
      "suit": "SPADES",
      "code": "5S"
    }
  ]
}
```
