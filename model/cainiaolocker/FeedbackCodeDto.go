package cainiaolocker

// FeedbackCodeDTO 
type FeedbackCodeDTO struct {
    // 异常反馈编码
    FeedbackCode   string `json:"feedback_code,omitempty" xml:"feedback_code,omitempty"`
    // 异常反馈编码描述
    FeedbackDesc   string `json:"feedback_desc,omitempty" xml:"feedback_desc,omitempty"`
}
