# Commentator

![Test](https://github.com/dannyh79/commentator/actions/workflows/ci.yml/badge.svg)

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

- [x] 基於短評功能，設計一套資料結構和後端服務，用於管理和顯示短評 -> See [/docs/data-structure.md](https://github.com/dannyh79/commentator/blob/main/docs/data-structure.md)
- [x] 基於上述設計，規劃數個 RESTful API 支援以下功能 -> See [/docs/rest-endpoints.md](https://github.com/dannyh79/commentator/blob/main/docs/rest-endpoints.md)
    - 新增短評
    - 獲取特定影片的所有短評
    - 對短評進行檢舉
    - 管理員對短評進行審核和處理
- [x] 請挑選其中一個 API ，進行 Pseudo Code 設計，完成這個 API 的功能：
    - 設計時請兼具可讀性、可測試性，以及效能考量 -> See [/docs/app-structure.md](https://github.com/dannyh79/commentator/blob/main/docs/app-structure.md)
    - 請設計 test case 來確認此 API 可行 -> See [/internal/rest/v1/routes_test.go](https://github.com/dannyh79/commentator/blob/main/internal/rest/v1/routes_test.go) for acceptance tests
- [x] 任何想要提出的創意或建議
    1. Press `.` to open web editor
    2. In editor's search section, use regex (click `.*` button) then search by `enhancement|gotcha`

## Gotchas

- Should be some log masking mechanism in place to mask any PII (Personally Identifiable Information)
