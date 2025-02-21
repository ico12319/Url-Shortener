package parse_url

import (
	"unicode/utf8"
	"url_shortener/random_string_generator"
	"url_shortener/shortened_url"
	"url_shortener/url_reader"
)

type ParseUrl struct {
	Shortened []shortened_url.ShortenedUrl
	reader    *url_reader.UrlReader
}

func NewParsers() *ParseUrl {
	return &ParseUrl{Shortened: make([]shortened_url.ShortenedUrl, 0), reader: url_reader.NewUrlReader()}
}

func (p *ParseUrl) Parse(fileName string) {
	p.reader.FillUrls(fileName)
	for i := 0; i < len(p.reader.Urls); i++ {
		originalUrlObject := p.reader.Urls[i]
		generator := random_string_generator.CreateGenerator(originalUrlObject.UrlAddress, utf8.RuneCountInString(originalUrlObject.UrlAddress))
		shortenedKey := generator.GenerateRandomString()
		shortenedObject := shortened_url.NewShortenedUrl(&originalUrlObject, shortenedKey)
		p.Shortened = append(p.Shortened, *shortenedObject)
	}
}
