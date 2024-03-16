package fulltextsearch

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int    `xml:"id"`
}

func LoadDocuments(dumpPath string) ([]Document, error) {
	f, err := os.Open(dumpPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	dec := xml.NewDecoder(r)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
