package miniappcloud

// File 
type File struct {
    // 文件地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 文件查询错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 文件查询返回结果，1-成功，0-失败
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 文件id
    FileId   string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}
