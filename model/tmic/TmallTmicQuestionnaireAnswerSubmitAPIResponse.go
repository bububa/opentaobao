package tmic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交问卷答案 API返回值 
tmall.tmic.questionnaire.answer.submit

天猫新品创新中心对外开放问卷，提交问卷答案
*/
type TmallTmicQuestionnaireAnswerSubmitAPIResponse struct {
    model.CommonResponse
    TmallTmicQuestionnaireAnswerSubmitAPIResponseModel
}

// 提交问卷答案 成功返回结果
type TmallTmicQuestionnaireAnswerSubmitAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_tmic_questionnaire_answer_submit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误提示语
    BizErrorInfo   string `json:"biz_error_info,omitempty" xml:"biz_error_info,omitempty"`
    // 错误编码
    BizErrorCode   string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
    // 是否调用成功
    BizSuccess   bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
