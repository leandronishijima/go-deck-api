# Go Deck API

API to handle and cards to be used in card games.

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
```bash
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

| Name    | Required | Description        | Default value | Example                                |
| ------- | -------- | ------------------ | ------------- | -------------------------------------- |
| deck id | required | the deck id (uuid) | -             | "2d4ee730-94cc-4585-b02d-47238aee59f9" |

- Example request
```bash
curl "http://localhost:8080/api/deck/open/${deck_id}"
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

| Name    | Required | Description                          | Default value | Example                                |
| ------- | -------- | ------------------------------------ | ------------- | -------------------------------------- |
| deck id | required | the deck id (uuid)                   | -             | "2d4ee730-94cc-4585-b02d-47238aee59f9" |
| count   | required | How many cards to draw from the deck | -             | 3                                      |

- Example request
```bash
curl -X "PATCH" "http://localhost:8080/api/deck/${deck_id}/draw" \
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

## Run locally

```sh
go run .
```

- The API will be available in the http://localhost:8080 

## Run tests

```sh
go test
```
