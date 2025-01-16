
package download

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strconv"
  "time"

  "xkcd/internal/types"
)

func PerformDownload() {
  // create data directory.
  fmt.Println("\nPrepare...")
  if err := os.MkdirAll(types.DATA_DIR, os.ModePerm); err != nil {
    log.Fatalf("Error creating directory: %v", err)
    return
  }
  // loop over all comic ids so far.
  for i := 1; i <= 3038; i++ {
    if i == 404 { // XKCD id 404 does not exist.
      i++
    }
    comic_id_str := strconv.Itoa(i)
    file_path := types.DATA_DIR + "/xkcd_" + comic_id_str + ".json"
    // check if data has already been saved to file.
    if ! fileExists(file_path) {
      // download data.
      fmt.Println("\nDownloading...")
      url := types.URL_START + comic_id_str + types.URL_END
      comic, err := fetchXKCDComic(url)
      if err != nil {
        log.Fatalf("Error fetching XKCD data: %v", err)
      }
      // show data.
      fmt.Printf(" - title:   %s\n", comic.Title)
      fmt.Printf(" - alt txt: %s\n", comic.Alt)
      fmt.Printf(" - img url: %s\n", comic.IMG)
      fmt.Println("\nSaving...")
      // save data.
      if err = saveToFile(file_path, comic); err != nil {
        log.Fatalf("Error saving to file: %v", err)
      }
      fmt.Printf("Download and save done.\n\n")
    }
  }
  fmt.Printf("Loop done.\n\n")
}

func fileExists(filePath string) bool {
  _, err := os.Stat(filePath)
  if os.IsNotExist(err) {
    return false
  }
  return err == nil
}

func fetchXKCDComic(url string) (*types.XKCDComic, error) {
  /// make the http request.
  client := http.Client{
    Timeout: 10 * time.Second,
  }
  resp, err := client.Get(url)
  if err != nil {
    return nil, fmt.Errorf("Error making GET request: %w", err)
  }
  defer resp.Body.Close()
  /// read the response body.
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, fmt.Errorf("Error reading response body: %w", err)
  }
  /// parse the json response.
  var comic types.XKCDComic
  if err := json.Unmarshal(body, &comic); err != nil {
    return nil, fmt.Errorf("Error unmarshalling JSON: %w", err)
  }
  return &comic, nil
}

func saveToFile(filename string, data *types.XKCDComic) error {
  // convert the comic data to json.
  jsonData, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    return fmt.Errorf("Error marshalling JSON: %w", err)
  }
  // write the json data to a file.
  err = os.WriteFile(filename, jsonData, 0644)
  if err != nil {
    return fmt.Errorf("Error writing to file: %w", err)
  }
  return nil
}

