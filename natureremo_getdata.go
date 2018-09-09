package main

import ( 
    "os"
    "net/http"
    "fmt"
    "time"
    "io/ioutil"
    "encoding/json"
)

var (
    url = "https://api.nature.global/1/devices"
    accesstoken = "Set accesstoken"
)


type NatureRemoJson struct {
    Id   string `json:"id"`
    Name string `json:"name"`
    Newest_events Newest_events `json:"newest_events"`
}

type Newest_events struct {
    Te Te `json:"te"`
    Il Il `json:"il"`
    Hu Hu `json:"hu"`
}

type Te struct {
    Created_at time.Time  `json:"created_at"`
    Val float64 `json:val`
}

type Il struct {
    Created_at time.Time  `json:"created_at"`
    Val float64 `json:val`
}

type Hu struct {
    Created_at time.Time  `json:"created_at"`
    Val float64 `json:val`
}




func main() {
    TickerNatureRemo()
}



func TickerNatureRemo(){

    ticker := time.NewTicker(600 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:

            natureremo_json := GetJsonFromNatureRemoAPI()
            WriteDataToText("temperature.txt", natureremo_json[0].Newest_events.Te.Created_at, natureremo_json[0].Newest_events.Te.Val)
            WriteDataToText("humidity.txt", natureremo_json[0].Newest_events.Hu.Created_at, natureremo_json[0].Newest_events.Hu.Val)
           
        }
    }
}





func GetJsonFromNatureRemoAPI() []NatureRemoJson {
    
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        panic(err)
    }

    req.Header.Add("Authorization", accesstoken)

    resp, err := http.DefaultClient.Do(req)
   
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        panic(err)
    }

    var natureremo_json []NatureRemoJson

    if err := json.Unmarshal([]byte(body), &natureremo_json);err != nil {
        fmt.Println(err)
    }

    return natureremo_json
}







func WriteDataToText(Filename string, Created_at time.Time, Val float64){

    file, err := os.OpenFile(Filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    
    if err != nil {
        panic(err)   
    }
    defer file.Close()

    //fix time zone
    jst := time.FixedZone("Asia/Tokyo", 9*60*60)
    Created_at = Created_at.In(jst)

    fmt.Fprintln(file, Created_at.Format(time.RFC3339), Val)
}



