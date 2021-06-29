package icbu

// CustomInfo 
type CustomInfo struct {
    // 定制内容
    CustomContents   []CustomContent `json:"custom_contents,omitempty" xml:"custom_contents>custom_content,omitempty"`
}
