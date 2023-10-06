package contracts

type Controller interface {
  exec() HttpResponse
}

