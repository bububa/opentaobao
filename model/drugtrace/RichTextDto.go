package drugtrace

// RichTextDTO 
type RichTextDTO struct {
    // 图片
    Pictures   []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
    // 文字
    Text   string `json:"text,omitempty" xml:"text,omitempty"`
}
