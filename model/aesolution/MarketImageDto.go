package aesolution

// MarketImageDto 
type MarketImageDto struct {
    // The image url needs to be obtained via uploading the image through Aliexpress API: aliexpress.photobank.redefining.uploadimageforsdk(https://developers.aliexpress.com/en/doc.htm?docId=30186&docType=2)
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // 1 represents 3:4 rectangle(resolution at least 750*1000) image while 2 represents 1:1 square image(Resolution at least 800*800).
    ImageType   string `json:"image_type,omitempty" xml:"image_type,omitempty"`
}
