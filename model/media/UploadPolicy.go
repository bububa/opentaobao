package media

// UploadPolicy 
type UploadPolicy struct {

    // 限制用户上传文件的大小。 若用户上传文件大小超过size_limit，无法上传成功。
    SizeLimit   int64 `json:"size_limit,omitempty"`

    // 限制用户上传文件的类型，多个值用；分隔。 可设置的值为：image/jpeg,image/png,image/webp等。 若用户上传文件的mime类型不在mime_limit范围内，无法上传成功。
    MimeLimit   string `json:"mime_limit,omitempty"`

}
