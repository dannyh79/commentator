# REST Endpoint Design

## Endpoints

All endpoints are versioned for consitency sake:

- `POST /v1/shows/{{ show id }}/comments` - 新增短評 (implemented)

- `GET /v1/shows/{{ show id }}/comments?filter=upvote:10&sort=upvote:desc&size=10&page=3` - 獲取特定影片的所有短評
  - Support (defined) querystrings here for the ease of modeling data for consumers
  - Could apply caching among layers here if consistency is not of much concern

- `POST /v1/shows/{{ show id }}/comments/{{ comment id }}:report` - 對短評進行檢舉

- `POST /v1/admin/shows/{{ show id }}/comments/{{ comment id }}:review` - 管理員對短評進行審核和處理
  - Namespace internal API with /admin

## Gotchas

- All endpoints should be protected by CORS policies and limit the use to authorized users/admins only
