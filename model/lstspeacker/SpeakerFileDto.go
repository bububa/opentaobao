package lstspeacker

// SpeakerFileDto 
type SpeakerFileDto struct {
    // md5
    Md5   string `json:"md5,omitempty" xml:"md5,omitempty"`
    // url
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // fileId
    FileId   string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}
