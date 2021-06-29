package tuanhotel

// ImagePathResultVoList 
type ImagePathResultVoList struct {
    // 在图片中心的地址路径
    ImagePath   string `json:"image_path,omitempty" xml:"image_path,omitempty"`
    // 异常信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 图片id
    ImageUid   string `json:"image_uid,omitempty" xml:"image_uid,omitempty"`
}
