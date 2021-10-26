# MuV 

MuV (read muvi) is a microservice prototype that incorporates OMDB API to provides movie info.

## Features

- Get Movie By Search Word of the Movie
- Get Movie Detail By IMDBID


## API Specification

#### Get Movie By certain search word

#### REST
```bash
  GET /movie
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `search` | `string` | **Required**. SearchWord for the movie|
| `p` | `string` | **Optional**. Page to visit (default to 1)|

#### Get Movie Detail By ImdbID

```bash
  GET /movie/detail
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `imdbID`      | `string` | **Required**. imdbID of the movie |

#### GRPC
- GRPC specification can refer to file [handler.proto](https://github.com/mfsyahrz/submission/blob/master/project/internal/transport/grpc/handler.proto) inside this project repository

## Tech Stack

**language:** [golang](https://golang.org/).

**database:**  [postgres](https://www.postgresql.org/).

**protocol:** [REST](https://restfulapi.net/), [GRPC](https://grpc.io/).

## How To Run
  - Please visit [How To Run](https://github.com/mfsyahrz/submission/blob/master/project/docs/how_to_run.md) which located in docs folder.

## Authors

- [@mfsyahrz](https://www.github.com/mfsyahrz)

## License

[MIT](https://choosealicense.com/licenses/mit/)
