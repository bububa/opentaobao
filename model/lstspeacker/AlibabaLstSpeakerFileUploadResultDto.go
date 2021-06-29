package lstspeacker

// AlibabaLstSpeakerFileUploadResultDto 
type AlibabaLstSpeakerFileUploadResultDto struct {

    // 执行结果
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    

    // SpeakerFileDto
    
    Module   *SpeakerFileDto `json:"module,omitempty" xml:"module,omitempty"`
    

    // 错误码
    
    ErroMessage   string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
    

    // 错误码
    
    ErroCode   string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
    

}
