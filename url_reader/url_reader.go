package url_reader

import (
	"bufio"
	"fmt"
	"os"
	"url_shortener/url"
)

const EMPTY_STRING = ""

type UrlReader struct {
	Urls []url.Url
}

func NewUrlReader() *UrlReader {
	return &UrlReader{make([]url.Url, 0)}
}

func (reader *UrlReader) FillUrls(fileName string) {
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		token := scanner.Text()
		newUrl := url.NewUrl(token)
		if newUrl != nil {
			reader.Urls = append(reader.Urls, *newUrl)
		} else {
			fmt.Println("Invalid token: ", token)
		}

	}
}
