package contracts

type Controller interface {
  Exec() HttpResponse
}

