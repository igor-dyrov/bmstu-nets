package main

import (

	"fmt" // пакет для форматированного ввода вывода

	"net/http" // пакет для поддержки HTTP протокола

	"strings" // пакет для работы с UTF-8 строками

	"log" // пакет для логирования

	"sort"

	"github.com/RealJK/rss-parser-go"

)

type rssItems []rss.Item

func (items rssItems) Less(i, j int) bool {
	return items[i].PubDate < items[j].PubDate 
}

func (items rssItems) Len() int {
	return len(items)
}

func (items rssItems) Swap(i, j int) { items[i], items[j] = items[j], items[i] }

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //анализ аргументов,

	fmt.Println(r.Form) // ввод информации о форме на стороне сервера

	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {

		fmt.Println("key:", k)

		fmt.Println("val:", strings.Join(v, ""))

	}

	fmt.Fprintf(w, "Test!") // отправляем данные на клиентскую сторону

}

func GetCommonRSS(w http.ResponseWriter, r *http.Request) {
	rss1, err := rss.ParseRSS("http://blagnews.ru/rss_vk.xml")
	rss2, err2 := rss.ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")
	rss3, err3 := rss.ParseRSS("https://lenta.ru/rss")

	fmt.Fprintf(w, `<a style="margin: 100px; font-size: 24px;" href="/rss1" class="rss-link">Blagnews</a>`)
	fmt.Fprintf(w, `<a style="margin: 100px; font-size: 24px;" href="/rss2" class="rss-link">Rss Board</a>`)
	fmt.Fprintf(w, `<a style="margin: 100px; font-size: 24px;" href="/rss3" class="rss-link">Lenta</a>`)

	if err != nil && err2 != nil && err3 != nil {
		rssList := append(rss1.Channel.Items, rss2.Channel.Items...)
		rssList = append(rssList, rss3.Channel.Items...)

		rssToSort := rssItems(rssList)

		sort.Sort(rssToSort)

		for v := range rssToSort {
			item := rssToSort[v]
			
			fmt.Println()

			fmt.Printf("Item Number : %d\n", v)
			fmt.Fprintf(w, fmt.Sprintf("<h1>Title: %s</h1>", item.Title))
			fmt.Fprintf(w, fmt.Sprintf("<h2>PubDate: %s</h2>", item.PubDate))

			fmt.Printf("Title : %s\n", item.Title)

			fmt.Printf("Link : %s\n", item.Link)

			fmt.Printf("Description : %s\n", item.Description)
			fmt.Fprintf(w, item.Description)

			fmt.Printf("Guid : %s\n", item.Guid.Value)
		}
	}
}

func GetFirstRSS(w http.ResponseWriter, r *http.Request) {
	rssObject, err := rss.ParseRSS("http://blagnews.ru/rss_vk.xml")

	fmt.Fprintf(w, "<h1>RSS Channel: Blag News</h1>")

	if err != nil {

		fmt.Printf("Title : %s\n", rssObject.Channel.Title)

		fmt.Printf("Generator : %s\n", rssObject.Channel.Generator)

		fmt.Printf("PubDate : %s\n", rssObject.Channel.PubDate)

		fmt.Printf("LastBuildDate : %s\n", rssObject.Channel.LastBuildDate)

		fmt.Printf("Description : %s\n", rssObject.Channel.Description)

		fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

		for v := range rssObject.Channel.Items {

			item := rssObject.Channel.Items[v]

			fmt.Println()

			fmt.Printf("Item Number : %d\n", v)
			fmt.Fprintf(w, fmt.Sprintf("<h2>Title: %s</h2>", item.Title))

			fmt.Printf("Title : %s\n", item.Title)

			fmt.Printf("Link : %s\n", item.Link)

			fmt.Printf("Description : %s\n", item.Description)
			fmt.Fprintf(w, item.Description);

			fmt.Printf("Guid : %s\n", item.Guid.Value)

		}

	}
}

func GetSecondRSS(w http.ResponseWriter, r *http.Request) {
	rssObject, err := rss.ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")

	fmt.Fprintf(w, "<h1>RSS Channel: RSS Board</h1>")

	if err != nil {

		fmt.Printf("Title : %s\n", rssObject.Channel.Title)

		fmt.Printf("Generator : %s\n", rssObject.Channel.Generator)

		fmt.Printf("PubDate : %s\n", rssObject.Channel.PubDate)

		fmt.Printf("LastBuildDate : %s\n", rssObject.Channel.LastBuildDate)

		fmt.Printf("Description : %s\n", rssObject.Channel.Description)

		fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

		for v := range rssObject.Channel.Items {

			item := rssObject.Channel.Items[v]

			fmt.Println()

			fmt.Printf("Item Number : %d\n", v)
			fmt.Fprintf(w, fmt.Sprintf("<h1>Title: %s</h1>", item.Title))

			fmt.Printf("Title : %s\n", item.Title)

			fmt.Printf("Link : %s\n", item.Link)

			fmt.Printf("Description : %s\n", item.Description)
			fmt.Fprintf(w, item.Description);

			fmt.Printf("Guid : %s\n", item.Guid.Value)

		}

	}
}

func GetThirdRSS(w http.ResponseWriter, r *http.Request) {
	rssObject, err := rss.ParseRSS("https://lenta.ru/rss")

	fmt.Fprintf(w, "<h1>RSS Channel: Lenta</h1>")

	if err != nil {

		fmt.Printf("Title : %s\n", rssObject.Channel.Title)

		fmt.Printf("Generator : %s\n", rssObject.Channel.Generator)

		fmt.Printf("PubDate : %s\n", rssObject.Channel.PubDate)

		fmt.Printf("LastBuildDate : %s\n", rssObject.Channel.LastBuildDate)

		fmt.Printf("Description : %s\n", rssObject.Channel.Description)

		fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

		for v := range rssObject.Channel.Items {

			item := rssObject.Channel.Items[v]

			fmt.Println()

			fmt.Printf("Item Number : %d\n", v)
			fmt.Fprintf(w, fmt.Sprintf("<h1>Title: %s</h1>", item.Title))

			fmt.Printf("Title : %s\n", item.Title)

			fmt.Printf("Link : %s\n", item.Link)

			fmt.Printf("Description : %s\n", item.Description)
			fmt.Fprintf(w, item.Description);

			fmt.Printf("Guid : %s\n", item.Guid.Value)

		}

	}
}

func main() {

	http.HandleFunc("/", GetCommonRSS) 
	http.HandleFunc("/rss1", GetFirstRSS)
	http.HandleFunc("/rss2", GetSecondRSS)
	http.HandleFunc("/rss3", GetThirdRSS)

	err := http.ListenAndServe(":9007", nil) // задаем слушать порт

	if err != nil {

		log.Fatal("ListenAndServe: ", err)

	}

}