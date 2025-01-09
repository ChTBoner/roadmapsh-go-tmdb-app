package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var api_prefix string = "https://api.themoviedb.org/3/movie"

type (
	Result struct {
		Title string
	}
	TmdbResp struct {
		Results []Result
	}
)

func main() {
	// parse command line args
	typePtr := flag.String("type", "top", "playing, popular, top or upcoming")
	flag.Parse()

	var mtype string
	switch *typePtr {
	case "top":
		mtype = "top_rated"
	case "popular":
		mtype = "popular"
	case "upcoming":
		mtype = "upcoming"
	case "playing":
		mtype = "now_playing"
	default:
		fmt.Println("Invalid option: Usage: top, playing, upcoming or popular")
		os.Exit(1)
	}

	tmdb_call(mtype)
}

func tmdb_call(mtype string) {
	token := os.Getenv("TMDB_TOKEN")
	url := fmt.Sprintf("%s/%s", api_prefix, mtype)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	bearer := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var f TmdbResp
	err = json.Unmarshal(data, &f)
	if err != nil {
		panic(err)
	}

	for _, v := range f.Results {
		fmt.Println(v.Title)
	}
}
