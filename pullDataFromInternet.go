package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

type SitemapIndex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (L Location) String() string {
  return fmt.Sprintf(L.Loc)
}
func main() {
  // sitemaps that conatian links to all catogrize site maps
  resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var s SitemapIndex
  xml.Unmarshal(bytes, &s)
  fmt.Println(s.Locations)
}
