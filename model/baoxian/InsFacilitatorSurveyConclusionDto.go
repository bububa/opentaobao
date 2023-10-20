package baoxian

import (
	"sync"
)

// InsFacilitatorSurveyConclusionDto 结构体
type InsFacilitatorSurveyConclusionDto struct {
	// 业务单号
	BizNo string `json:"biz_no,omitempty" xml:"biz_no,omitempty"`
	// 物流单号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 勘察附件
	SurveyAttachments string `json:"survey_attachments,omitempty" xml:"survey_attachments,omitempty"`
	// 勘察结论
	SurveyConclusion string `json:"survey_conclusion,omitempty" xml:"survey_conclusion,omitempty"`
	// 勘察结论描述
	SurveyConclusionDesc string `json:"survey_conclusion_desc,omitempty" xml:"survey_conclusion_desc,omitempty"`
	// 扩展参数
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
}

var poolInsFacilitatorSurveyConclusionDto = sync.Pool{
	New: func() any {
		return new(InsFacilitatorSurveyConclusionDto)
	},
}

// GetInsFacilitatorSurveyConclusionDto() 从对象池中获取InsFacilitatorSurveyConclusionDto
func GetInsFacilitatorSurveyConclusionDto() *InsFacilitatorSurveyConclusionDto {
	return poolInsFacilitatorSurveyConclusionDto.Get().(*InsFacilitatorSurveyConclusionDto)
}

// ReleaseInsFacilitatorSurveyConclusionDto 释放InsFacilitatorSurveyConclusionDto
func ReleaseInsFacilitatorSurveyConclusionDto(v *InsFacilitatorSurveyConclusionDto) {
	v.BizNo = ""
	v.LogisticsNo = ""
	v.SurveyAttachments = ""
	v.SurveyConclusion = ""
	v.SurveyConclusionDesc = ""
	v.ExtendInfo = ""
	poolInsFacilitatorSurveyConclusionDto.Put(v)
}
