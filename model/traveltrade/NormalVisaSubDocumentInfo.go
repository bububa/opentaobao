package traveltrade

// NormalVisaSubDocumentInfo 
type NormalVisaSubDocumentInfo struct {
    // 子材料文档编号，10001:护照封面，10002:护照首页，11401:去程机票，11402:返程机票
    DocType   int64 `json:"doc_type,omitempty" xml:"doc_type,omitempty"`
    // 文件类型
    FileType   string `json:"file_type,omitempty" xml:"file_type,omitempty"`
    // base64编码的文件内容
    FileContent   string `json:"file_content,omitempty" xml:"file_content,omitempty"`
}
