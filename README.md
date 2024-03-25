# Commentator

An example app with REST API `POST /v1/shows/:id/comments`. See /docs for more.

## Getting Started

```shell
# go version 1.22.1 assumed
go mod download && go run main.go

curl -X POST localhost:8080/v1/shows/1/comments -d '{"user_id":2,"comment":"some comments"}'
# HTTP/1.1 201 Created
# Content-Type: application/json; charset=utf-8
# Date: Mon, 25 Mar 2024 11:48:48 GMT
# Content-Length: 67

# {"result":{"user_id":2,"comment":"some commentsssssss","upvote":0}}%
```

## Todos

- [ ] 基於短評功能，設計一套資料結構和後端服務，用於管理和顯示短評
- [ ] 基於上述設計，規劃數個 RESTful API 支援以下功能
    - 新增短評
    - 獲取特定影片的所有短評
    - 對短評進行檢舉
    - 管理員對短評進行審核和處理
- [ ] 請挑選其中一個 API ，信行 Pseudo Code 設計，晚成這個 API 的功能：
    - 設計時請兼具可讀性、可測試性，以及效能考量
    - 請設計test case 來確認此 API 可行
- [ ] 任何想要提出的創意或建議

## Gotchas

- Should be some log masking mechanism in place to mask any PII (Personally Identifiable Information)
