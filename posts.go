package main

import (
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/errors"
	"gopkg.in/headzoo/surf.v1"
)

// PostMeta - Meta Data for posts
type PostMeta struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// Post - Post content data
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
	Slug  string `json:"slug"`
}

var allPostsURL = "https://www.buymeacoffee.com/" + os.Getenv("BMC_USERNAME") + "/posts"
var singlePostURL = "https://www.buymeacoffee.com/" + os.Getenv("BMC_USERNAME")

// GetAllPosts - get all posts for provided username
func GetAllPosts() []PostMeta {
	bow := surf.NewBrowser()
	err := bow.Open(allPostsURL)
	if err != nil {
		panic(err)
	}

	var posts []PostMeta

	container := bow.Find("div.post-title-wrap")

	container.Each(func(index int, item *goquery.Selection) {
		postLink, exists := item.Find("a").Attr("href")
		postTitle, error := item.Find("div.post-title").Html()

		if !exists {
			return
		}

		if error != nil {
			return
		}

		posts = append(posts, PostMeta{
			Title: postTitle,
			Slug:  postLink[len(singlePostURL+"/"):],
		})
	})

	return posts
}

// GetPostBySlug - get post details for specified url
func GetPostBySlug(slug string) (Post, error) {
	var post = Post{
		Slug: slug,
	}
	bow := surf.NewBrowser()
	err := bow.Open(singlePostURL + "/" + slug)
	if err != nil {
		panic(err)
	}
	postTitle, err := bow.Find("a > div.post-title-s").Html()

	if err != nil {
		return post, errors.New("Couldn't get title")
	}

	postBody, err := bow.Find("div.post-content>.mBox").Html()

	if err != nil {
		return post, errors.New("Couldn't get post body")
	}

	postTime, err := bow.Find("div.post-single-content div > p.text-fs-16.clr-grey.halvetica-light.mg-b-0.xs-fs-14").Html()

	if err != nil {
		return post, errors.New("Couldn't get post time")
	}

	post.Body = postBody
	post.Title = postTitle
	post.Date = postTime

	return post, nil
}
