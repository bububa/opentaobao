package product

// UploadImageResp 
/* model for simplify = false
type UploadImageResp struct {

    // 返回的图片地址
    
    Url   string `json:"url,omitempty"`
    

    // 图片哈希值
    
    Hash   string `json:"hash,omitempty"`
    

    // 图片高度
    
    Height   int64 `json:"height,omitempty"`
    

    // 图片宽度
    
    Width   int64 `json:"width,omitempty"`
    

}
*/

// UploadImageResp 
type UploadImageResp struct {

    // 返回的图片地址
    Url   string `json:"url,omitempty"`

    // 图片哈希值
    Hash   string `json:"hash,omitempty"`

    // 图片高度
    Height   int64 `json:"height,omitempty"`

    // 图片宽度
    Width   int64 `json:"width,omitempty"`

}
