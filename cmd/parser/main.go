package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	handlerBd "parser_habr/internal"
)

/*
https://habr.com/ru/post/493088/
https://github.com/noorsoft-mobile/go-cinema-parser/blob/master/main.go
https://github.com/geziyor/geziyor

//DONE Узнать что такое амперсант / звездочка
https://habr.com/ru/post/339192/

*/

func main() {
	handlerParserToJson()
	//TODO разобраться как разделять на модули
	//TODO Подключить к базе данных + laravel + логирование
	handlerBd.Get()
	//TODO Подключить к боту
	//TODO Сделать меню выбора тегов и фильтрации
	//TODO разобраться как разместить на сервере
	//TODO CI/CD github
}

func handlerParserToJson() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://habr.com/ru/all/"},
		ParseFunc: parserHabr,
		//TODO разобраться как генерируется результат
		//TODO как посмотреть подобное исходный код
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func parserHabr(g *geziyor.Geziyor, r *client.Response) {
	//tm-article-snippet__title-link
	r.HTMLDoc.Find(".tm-article-snippet__title-link").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("span").Text()
		//fmt.Println(title)
		//TODO как происходит выгрузка
		//TODO прочитать про map interface
		g.Exports <- map[string]interface{}{
			"title": title,
			//TODO Дозаполнить данными
		}
	})
}
