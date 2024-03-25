# App Structure

![](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

> PC: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

The app is designed to be layered in the following order, from outer to inner:

```
Entry point (/main.go)
  -> REST routes (/internal/rest) with domain usecases
    -> Domain Usecases with repositories (e.g., /internal/shows/show_comments.go)
      -> Domain entities (e.g., /internal/shows/entities/comment.go)

  -> Repositories (/internal/repo)
```

Such design has the following benefits:
- Minimal setup to unit test modules; dependencies can be injected by mocks
- Plugable persistent storage; usecases code on repository interface, not actual implementations
- ...
