package lstvending

// VendingImageDTO 
type VendingImageDTO struct {
    // 图片唯一标识
    ImgPathId   string `json:"img_path_id,omitempty" xml:"img_path_id,omitempty"`
    // 图片访问地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
}
