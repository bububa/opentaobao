package eticket

// UploadImgCallbackResp 
type UploadImgCallbackResp struct {
    // 扩展属性
    AttributeMap   string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
    // 图片在淘宝的文件名
    FileName   string `json:"file_name,omitempty" xml:"file_name,omitempty"`
}
