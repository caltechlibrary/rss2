//
// rss2 is a golang package for working with RSS 2 feeds and documents.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2018, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package rss2

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

const Version = `v0.0.4`

type CustomAttrs []xml.Attr

type RSS2 struct {
	XMLName xml.Name `xml:"rss" json:"-"`
	Version string   `xml:"version,attr" json:"version"`

	// Required
	Title       string `xml:"channel>title" json:"title"`
	Link        string `xml:"channel>link" json:"link"`
	Description string `xml:"channel>description" json:"description"`

	// Optional
	Language       string `xml:"channel>language,omitempty" json:"language,omitempty"`
	Copyright      string `xml:"channel>copyright,omitempty" json:"copyright,omitempty"`
	ManagingEditor string `xml:"channel>managingEditor,omitempty" json:"managingEditor,omitempty"`
	WebMaster      string `xml:"channel>webMaster,omitempty" json:"webMaster,omitempty"`
	PubDate        string `xml:"channel>pubDate,omitempty" json:"pubDate,omitempty"`
	LastBuildDate  string `xml:"channel>lastBuildDate,omitempty" json:"lastBuildDate,omitempty"`
	Category       string `xml:"channel>category,omitempty" json:"category,omitempty"`
	Generator      string `xml:"channel>generator,omitempty" json:"generator,omitempty"`
	Docs           string `xml:"channel>docs,omitempty" json:"docs,omitempty"`
	Cloud          string `xml:"channel>cloud,omitempty" json:"cloud,omitempty"`
	TTL            string `xml:"channel>ttl,omitempty" json:"ttl,omitempty"`
	Image          string `xml:"channel>image,omitempty" json:"image,omitempty"`
	Rating         string `xml:"channel>rating,omitempty" json:"rating,omitempty"`
	SkipHours      string `xml:"channel>skipHours,omitempty" json:"skipHours,omitempty"`
	SkipDays       string `xml:"channel>skipDays,omitempty" json:"skipDays,omitempty"`
	ItemList       []Item `xml:"channel>item,omitempty" json:"item,omitempty"`
}

type Item struct {
	// Optional according to Dave Winer
	Title string `xml:"title" json:"title,omitempty"`

	// Required
	Link string `xml:"link" json:"link"`

	// Optional
	Description string      `xml:"description,omitempty" json:"description,omitempty"`
	Author      string      `xml:"author,omitempty" json:"author,omitempty"`
	Category    string      `xml:"category,omitempty" json:"category,omitempty"`
	Content     string      `xml:"encoded,omitempty" json:"encoded,omitempty"`
	PubDate     string      `xml:"pubDate,omitempty" json:"pubDate,omitempty"`
	Comments    string      `xml:"comments,omitempty" json:"comments,omitempty"`
	Enclosure   string      `xml:"enclosure,omitempty" json:"enclosure,omitempty"`
	GUID        string      `xml:"guid,omitempty" json:"guid,omitempty"`
	Source      string      `xml:"source,omitempty" json:"source,omitempty"`
	OtherAttr   CustomAttrs `xml:",any,attr" json:"other_attrs,omitempty"`
}

// MarshalJSON() marshals the custom attributes that might
// be included in an RSS feed.
func (cattr CustomAttrs) MarshalJSON() ([]byte, error) {
	m := map[string]string{}
	for _, attr := range cattr {
		k := attr.Name.Local
		v := attr.Value
		if k != "" {
			m[k] = v
		}
	}
	return json.Marshal(m)
}

// Parse return an RSS2 document as a RSS2 structure.
func Parse(buf []byte) (*RSS2, error) {
	data := new(RSS2)
	err := xml.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RSS2) channel(dataPath string) (map[string]interface{}, error) {
	results := make(map[string]interface{})
	switch {
	case strings.Compare(dataPath, ".channel") == 0:
		// package and return all the channel fields
		results[".title"] = r.Title
		results[".link"] = r.Link
		results[".description"] = r.Description
		if r.PubDate != "" {
			results[".pubDate"] = r.PubDate
		}
	case strings.HasSuffix(dataPath, ".title"):
		results[".title"] = r.Title
	case strings.HasSuffix(dataPath, ".link"):
		results[".link"] = r.Link
	case strings.HasSuffix(dataPath, ".description"):
		results[".description"] = r.Description
	case strings.HasSuffix(dataPath, ".pubDate"):
		results[".pubDate"] = r.PubDate
	default:
		return nil, fmt.Errorf("Unknown data path %s", dataPath)
	}
	return results, nil
}

type rangeExpression struct {
	first int
	last  int
}

func getRange(listLength int, exp string) *rangeExpression {
	rexp := new(rangeExpression)
	rexp.first = 0
	rexp.last = listLength - 1

	if strings.Contains(exp, "-") == true {
		nums := strings.SplitN(exp, "-", 2)
		i, err := strconv.Atoi(nums[0])
		if err == nil {
			rexp.first = i
		}
		i, err = strconv.Atoi(nums[1])
		if err == nil {
			rexp.last = i
		}
	} else {
		i, err := strconv.Atoi(exp)
		if err == nil {
			rexp.first = i
			rexp.last = i
		}
	}
	return rexp
}

func (rexp *rangeExpression) inRange(val int) bool {
	if val >= rexp.first && val <= rexp.last {
		return true
	}
	return false
}

func (r *RSS2) items(dataPath string) (map[string]interface{}, error) {
	rexp := new(rangeExpression)
	rexp.first = 0
	rexp.last = len(r.ItemList) - 1

	// Get the range expression so we know when to add it to results.
	s := strings.Index(dataPath, "[")
	e := strings.Index(dataPath, "]")
	if s >= 0 && e >= 0 {
		rexp = getRange(len(r.ItemList), dataPath[s:e])
	}

	results := make(map[string]interface{})
	switch {
	case strings.HasSuffix(dataPath, ".title") == true:
		vals := []string{}
		for i, item := range r.ItemList {
			if rexp.inRange(i) == true {
				vals = append(vals, item.Title)
			}
		}
		results["title"] = vals
	case strings.HasSuffix(dataPath, ".link") == true:
		vals := []string{}
		for i, item := range r.ItemList {
			if rexp.inRange(i) == true {
				vals = append(vals, item.Link)
			}
		}
		results["link"] = vals
	case strings.HasSuffix(dataPath, ".description") == true:
		vals := []string{}
		for i, item := range r.ItemList {
			if rexp.inRange(i) == true {
				vals = append(vals, item.Description)
			}
		}
		results["description"] = vals
	case strings.HasSuffix(dataPath, ".content") == true:
		vals := []string{}
		for i, item := range r.ItemList {
			if rexp.inRange(i) == true {
				vals = append(vals, item.Content)
			}
		}
		results["content"] = vals
	case strings.HasSuffix(dataPath, ".pubDate") == true:
		vals := []string{}
		for i, item := range r.ItemList {
			if rexp.inRange(i) == true {
				vals = append(vals, item.PubDate)
			}
		}
		results["pubDate"] = vals
	case strings.HasSuffix(dataPath, ".comments") == true:
		vals := []string{}
		for i, item := range r.ItemList {
			if rexp.inRange(i) == true {
				vals = append(vals, item.Comments)
			}
		}
		results["comments"] = vals
	}
	return results, nil
}

// Filter given an RSS2 document return all the entries matching so we
// can apply return each of the data paths requested.
// e.g. .version, .channel.title, .channel.link, .item[].link,
// .item[].guid, .item[].title, .item[].description
func (r *RSS2) Filter(dataPaths []string) (map[string]interface{}, error) {
	var (
		err  error
		data map[string]interface{}
	)
	result := make(map[string]interface{})
	for _, dataPath := range dataPaths {
		switch {
		case strings.Compare(dataPath, ".version") == 0:
			result["version"] = r.Version
		case strings.HasPrefix(dataPath, ".channel"):
			data, err = r.channel(dataPath)
			// Merge data into results keyed' by path
			for _, val := range data {
				result[dataPath] = val
			}
		case strings.HasPrefix(dataPath, ".item[]"):
			data, err = r.items(dataPath)
			// Merge data into results keyed' by path
			for _, val := range data {
				result[dataPath] = val
			}
		default:
			return nil, fmt.Errorf("path %q not found", dataPath)
		}
	}
	if result == nil {
		return nil, fmt.Errorf("No data paths found")
	}
	return result, err
}
