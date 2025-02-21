package url

type Url struct {
	UrlAddress string
}

func NewUrl(address string) *Url {
	return &Url{UrlAddress: address}
}
