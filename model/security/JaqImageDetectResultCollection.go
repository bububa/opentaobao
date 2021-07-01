package security

// JaqImageDetectResultCollection 结构体
type JaqImageDetectResultCollection struct {
	// 响应消息结构体
	JaqImageDetectResultList []JaqImageDetectResult `json:"jaq_image_detect_result_list,omitempty" xml:"jaq_image_detect_result_list>jaq_image_detect_result,omitempty"`
}
