package url_converter

import (
	"bufio"
	"fmt"
	"os"
	"url_shortener/parse_url"
)

type UrlConverter struct {
	Parser *parse_url.ParseUrl
}

func NewConverter() *UrlConverter {
	return &UrlConverter{Parser: parse_url.NewParsers()}
}

func (converter *UrlConverter) WriteConvertedUrl(fromFileName string, toFileName string) {
	converter.Parser.Parse(fromFileName)
	outFile, err := os.Create(toFileName)
	if err != nil {
		fmt.Println("Error when trying to create the file")
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for i := 0; i < len(converter.Parser.Shortened); i++ {
		fmt.Fprintf(writer, "%s ", converter.Parser.Shortened[i].ShortKey)
		fmt.Fprintf(writer, "%s\n", converter.Parser.Shortened[i].CUrl.UrlAddress)
	}
	if err := writer.Flush(); err != nil {
		fmt.Println("Error when flushing the buffer")
		return
	}
}
