package lstspeacker

// AlibabaLstSpeakerStatusGetResultDto 
type AlibabaLstSpeakerStatusGetResultDto struct {
    // 执行结果
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 状态对象
    Module   *SpeakerOnLineStatus `json:"module,omitempty" xml:"module,omitempty"`
    // 错误码
    ErroMessage   string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
    // 错误码
    ErroCode   string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
}