package tmallgenie

// Author 
type Author struct {
    // 内容作者，对应音乐为作词、作曲人，对应小说故事为原著作者
    ContentAuthor   string `json:"content_author,omitempty" xml:"content_author,omitempty"`
    // 主播，演唱者，演播者
    VoiceAuthor   string `json:"voice_author,omitempty" xml:"voice_author,omitempty"`
}
