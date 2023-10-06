package contracts

type HttpResponse struct {
  StatusCode int `json:"statusCode"`
  Data interface{} `json:"data"`
}

func Ok(data interface{}) HttpResponse {
  return HttpResponse{StatusCode: 200, Data: data}
}

func Created(data interface{}) HttpResponse {
  return HttpResponse{StatusCode: 201, Data: data}
}

func NoContent() HttpResponse {
  return HttpResponse{StatusCode: 204}
}

func BadRequest(data interface{}) HttpResponse {
  return HttpResponse{StatusCode: 400, Data: data}
}

func NotFound(data interface{}) HttpResponse {
  return HttpResponse{StatusCode: 404, Data: data}
}

func ServerError(data interface{}) HttpResponse {
  return HttpResponse{StatusCode: 500, Data: data}
}

