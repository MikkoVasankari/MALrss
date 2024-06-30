package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Guid        string   `xml:"guid"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
}

func main() {
	arg := os.Args

	if len(arg) < 2 {
		fmt.Println("To use the application run it like: ")
		fmt.Println("anime for anime list | manga for manga list")
		fmt.Println("Example: " + arg[0] + " anime username")
		return
	}

	if arg[1] == "help" || arg[1] == "-h" || arg[1] == "--help" {
		fmt.Println("To use the application run it like: ")
		fmt.Println("anime for anime list | manga for manga list")
		fmt.Println("Example: " + arg[0] + " anime username")
		return
	}

	if arg[1] != "anime" && arg[1] != "manga" {
		fmt.Println("Please add argument for list type")
		fmt.Println("anime for anime list | manga for manga list")
		fmt.Println("Example: " + arg[0] + " anime username")
		return
	}

	var listType string

	switch arg[1] {
	case "anime":
		listType = "rw"
	case "manga":
		listType = "rm"
	default:
		fmt.Println("No valid list type was selected")
	}

	if len(arg) < 3 {
		fmt.Println("Give username as argument.")
		fmt.Println("Like: " + arg[0] + " 1 username")
		return
	} else if len(arg) > 3 {
		fmt.Println("Too many arguments.")
		return
	}

	user := arg[2]

	resp, err := http.Get("https://myanimelist.net/rss.php?type=" + listType + "&u=" + user)
	if err != nil {
		fmt.Printf("Get request failed: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Reading response body failed: %s", err)
		return
	}

	var rss RSS

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Printf("Unmarshaling XML failed: %s", err)
		return
	}

	fmt.Printf("Profile Title: %s\n", rss.Channel.Title)
	fmt.Printf("Profile Link: %s\n", rss.Channel.Link)
	fmt.Printf("Profile Description: %s\n", rss.Channel.Description)
	fmt.Printf("\n")
	for index, item := range rss.Channel.Items {
		fmt.Printf("%v: \n", index+1)
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Guid: %s\n", item.Guid)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Item PubDate: %s\n", item.PubDate)
		fmt.Printf("\n \n")
	}
}
