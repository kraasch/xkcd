
package search

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "path/filepath"
  "strings"

  "xkcd/internal/types"
)

func PerformSearch() {
  // interpret arguments as search terms.
  if len(os.Args) <= 2 {
    fmt.Println("Please provide search terms.")
    fmt.Println("Usage: xkcd search <search-terms>")
    return
  }
  terms := os.Args[2:]

  fmt.Println("Reading...")
  // directory containing the comic json files.
  comics, err := readComicsFromDir(types.DATA_DIR)
  if err != nil {
    log.Fatalf("Error reading comics: %v", err)
  }

  // search for terms and print result.
  fmt.Println("Search for:")
  for _, term := range terms {
    fmt.Printf(" - %s\n", term)
  }
  fmt.Println()
  fmt.Println("Found:")
  matches := Search(comics, terms)
  for _, comic := range matches {
    fmt.Printf(" - title: %s, image: %s\n", comic.Title, comic.IMG)
  }
  fmt.Printf("Done.\n\n")
}

// reads all the comic json files in a directory and unmarshals them into a slice.
func readComicsFromDir(dir string) ([]types.XKCDComic, error) {
  var comics []types.XKCDComic
  // read all files in the directory.
  files, err := os.ReadDir(dir)
  if err != nil {
    return nil, fmt.Errorf("Error reading directory: %w", err)
  }
  // iterate over all files in the directory.
  for _, file := range files {
    // only process json files.
    if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
      continue
    }
    // read the file content.
    filePath := filepath.Join(dir, file.Name())
    comicData, err := ioutil.ReadFile(filePath)
    if err != nil {
      return nil, fmt.Errorf("Error reading file %s: %w", filePath, err)
    }
    // unmarshal json data into xkcdcomic struct.
    var comic types.XKCDComic
    if err := json.Unmarshal(comicData, &comic); err != nil {
      return nil, fmt.Errorf("Error unmarshalling JSON from file %s: %w", filePath, err)
    }
    // add the comic to the slice.
    comics = append(comics, comic)
  }
  return comics, nil
}

// search searches through slice based on a query.
func Search(comics []types.XKCDComic, queries []string) []types.XKCDComic {
  var results []types.XKCDComic
  for i := range queries {
    queries[i] = strings.ToLower(queries[i])
  }
  for _, comic := range comics {
    for _, query := range queries {
      // check if the query is a substring of title or alt.
      if strings.Contains(strings.ToLower(comic.Title), query) || strings.Contains(strings.ToLower(comic.Alt), query) {
        results = append(results, comic)
      }
    }
  }
  return results
}

