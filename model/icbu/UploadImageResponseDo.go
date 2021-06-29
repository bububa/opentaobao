package icbu

// UploadImageResponseDo 
type UploadImageResponseDo struct {
    // 生成的图片名称
    FileName   string `json:"file_name,omitempty" xml:"file_name,omitempty"`
    // 生成的图片全路径URL
    PhotobankUrl   string `json:"photobank_url,omitempty" xml:"photobank_url,omitempty"`
    // 图片的唯一识别id
    FileId   int64 `json:"file_id,omitempty" xml:"file_id,omitempty"`
}
