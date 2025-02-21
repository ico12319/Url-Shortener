package shortened_url

import "url_shortener/url"

type ShortenedUrl struct {
	CUrl     *url.Url
	ShortKey string
}

func NewShortenedUrl(cUrl *url.Url, shortKey string) *ShortenedUrl {
	return &ShortenedUrl{CUrl: cUrl, ShortKey: shortKey}
}
