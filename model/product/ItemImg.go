package product

// ItemImg 
type ItemImg struct {
    // 图片链接地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 图片放在第几张（多图时可设置）
    Position   int64 `json:"position,omitempty" xml:"position,omitempty"`
}
