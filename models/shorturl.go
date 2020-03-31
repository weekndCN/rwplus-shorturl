package models

import (
    "fmt"
    "github.com/weekndCN/shorturl/pkg/shorturl"
)

type Shorturl struct {
    ID int `gorm: "primary_key" json:"id"`
    Shorten string `json:"shorten"`
    Lengthen string `json:"lengthen"`
}

func AddUrl(long_url string) (int, string, string) {
    shortenurl := &Shorturl{Lengthen : long_url}
    db.Create(&shortenurl)
    // new encoder
    encoder := shorturl.NewUrlEncoder(&shorturl.EncoderConfig{})
    // shorten设置uint64为了获取更大的长度，这里需要做显式转换
    shorten := encoder.EncodeURL(uint64(shortenurl.ID))
    // 更新shorten。生成shorten
    db.Model(&shortenurl).Update("shorten", shorten)
    return shortenurl.ID, shortenurl.Lengthen, shortenurl.Shorten
}


func GetLengthen(shorten string) (int, string) {
    var shortenurl Shorturl
    db.Select("id,lengthen").Where("shorten = ?", shorten).First(&shortenurl)
    if shortenurl.ID > 0 {
        return shortenurl.ID, shortenurl.Lengthen
    }
    return 0, ""
}


func GetShorten(lengthen string) (int, string) {
    var shortenurl Shorturl
    fmt.Println(lengthen)
    db.Select("id,shorten").Where("lengthen = ?", lengthen).First(&shortenurl)
    if shortenurl.ID > 0 {
        return shortenurl.ID, shortenurl.Shorten
    }
    return 0, ""
}
