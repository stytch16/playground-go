package NameData

import (
	"fmt"
	"strings"

	"github.com/stytch16/playground-go/BehindTheName/ParseHTML"
	"golang.org/x/net/html"
)

/* Process acquires data of name from html of behindthename.com/name/<name> */
func Process(theName string) (int, error) {
	/* Format name : 1st letter uppercase, rest lowercase */
	Name.theName = strings.ToUpper(string(theName[0])) + strings.ToLower(string(theName[1:]))

	/* Grab HTML data */
	htmlData, err := ParseHTML.FetchAndParse(URL + Name.theName)
	if err != nil {
		return 0, fmt.Errorf("process(%s) : %v\n", Name.theName, err)
	}

	/* Extract text from HTML */
	var htmlContent []string
	var contents func(*html.Node)
	contents = func(n *html.Node) {
		if n.Type == html.TextNode {
			s := strings.TrimSpace(n.Data)
			if s != "" {
				htmlContent = append(htmlContent, s)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
				/* do not descend into <script> or <style> elements. */
				break
			} else {
				contents(c)
			}
		}
	}
	contents(htmlData) /* This will populate htmlContent to become an array of text */

	/* Examine text to see if Name is not found or there exist multiple entries for that name. */
	if index("Name not found.", htmlContent) != -1 || index("There was no name definition found for "+Name.theName+".", htmlContent) != -1 {
		/* send signal that name is not found in behindthename.com */
		return 1, nil
	} else if index("There are multiple entries for "+Name.theName+".", htmlContent) != -1 {
		/* use the HTML data for 1st entry of Name */
		htmlData, err = ParseHTML.FetchAndParse(URL + Name.theName + "-1")
		if err != nil {
			return 0, fmt.Errorf("process(%s) : %v\n", Name.theName+"-1", err)
		}
		htmlContent = htmlContent[:0]
		contents(htmlData)
	}

	/* Targets are text to look for in htmlContent */
	targets := []string{
		"GENDER:",
		"USAGE:",
		"Meaning & History",
	}

	/* populateStruct populates the fields of struct Name based on htmlContent */
	populateStruct(targets[:], htmlContent[:])
	return 0, nil
}

/* populateStruct populates elements of a global struct.
The function invariant is that each of the strings desired from 'content'
will follow its corresponding text element in 'targets', e.g. if "GENDER:"
in 'targets' and "GENDER:" exists in 'content', then 'content' should look
something like [... ,"GENDER:", "Feminine", ... ] so "Feminine" is the desired string.
   However the target element 'Meaning & History' is a very special case because
its text has been seperated to multiple string elements due to non-text elements in the
middle. For that we must join strings and we take advantage of the fact that a period
should exist in every sentence of Meaning & History. See the function composeMeaningAndHist
for further details. */
func populateStruct(targets []string, content []string) {

	/* Function for joining strings that make up Meaning & History */
	composeMeaningAndHist := func(content []string) (output string) {
		var read []string /* buffer variable */
		for _, s := range content {
			/* Let Related Names section be our exit point since that is most common in behindthename name webpages */
			if s == "Related Names" {
				break
			}
			read = append(read, s)
			if strings.Index(s, ".") != -1 {
				output += strings.Join(read, " ") /* Join all text read and store it in output */
				read = read[:0]                   /* resets read string array */
			}
		}
		return output
	}

	for i, text := range content {
		if index(text, targets) != -1 {
			switch text {
			case "GENDER:":
				Name.gender = content[i+1] /* see Invariant */
			case "USAGE:":
				Name.usage = content[i+1]
			case "Meaning & History":
				Name.meaning = composeMeaningAndHist(content[i+1:])
			}
		}
	}
}

/* index checks to see if 'target' exists in 'strings' and returns index in 'strings'.
-1 otherwise. */
func index(target string, strings []string) int {
	for i, s := range strings {
		if s == target {
			return i
		}
	}
	return -1
}
