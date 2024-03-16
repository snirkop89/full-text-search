package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	fulltextsearch "github.com/snirkop89/full-text-search-engine"
)

func main() {
	f, err := os.Create("/tmp/indexer_cpu")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enenwiki-latest-abstract1.xml.gz", "dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	// Preparing
	fmt.Println("Full text search is in progress")
	start := time.Now()
	docs, err := fulltextsearch.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded %d documents in %s\n", len(docs), time.Since(start))

	// Indexing
	start = time.Now()
	idx := make(fulltextsearch.Index)
	idx.Add(docs)
	fmt.Printf("Indexed %d documents in %s\n", len(docs), time.Since(start))

	// Searching
	start = time.Now()
	matchedIDs := idx.Search(query)
	fmt.Printf("Search found %d documents in %s", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		fmt.Printf("%d\t%s\n", id, doc.Text)
	}
}
