package tmallgenie

// BatchContent 
type BatchContent struct {
    // 类目ID，具体参见开放平台类目相关描述
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    // 内容信息
    OpenContents   []OpenContent `json:"open_contents,omitempty" xml:"open_contents>open_content,omitempty"`
}
