package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click

// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

import "url_shortener/url_converter"

func main() {
	converter := url_converter.NewConverter()
	converter.WriteConvertedUrl("urls.txt", "shortened_urls_txt")
}
