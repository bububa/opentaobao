package product

// ItemImg 
/* model for simplify = false
type ItemImg struct {

    // 图片链接地址
    
    Url   string `json:"url,omitempty"`
    

    // 图片放在第几张（多图时可设置）
    
    Position   int64 `json:"position,omitempty"`
    

}
*/

// ItemImg 
type ItemImg struct {

    // 图片链接地址
    Url   string `json:"url,omitempty"`

    // 图片放在第几张（多图时可设置）
    Position   int64 `json:"position,omitempty"`

}
