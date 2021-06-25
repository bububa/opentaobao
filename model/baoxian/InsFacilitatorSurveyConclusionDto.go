package baoxian

// InsFacilitatorSurveyConclusionDto 
type InsFacilitatorSurveyConclusionDto struct {

    // 业务单号
    BizNo   string `json:"biz_no,omitempty"`

    // 物流单号
    LogisticsNo   string `json:"logistics_no,omitempty"`

    // 勘察附件
    SurveyAttachments   string `json:"survey_attachments,omitempty"`

    // 勘察结论
    SurveyConclusion   string `json:"survey_conclusion,omitempty"`

    // 勘察结论描述
    SurveyConclusionDesc   string `json:"survey_conclusion_desc,omitempty"`

    // 扩展参数
    ExtendInfo   string `json:"extend_info,omitempty"`

}
