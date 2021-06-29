package tmic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交单题答案 APIResponse
tmall.tmic.questionnaire.answer.push

问卷单题回答的提交
*/
type TmallTmicQuestionnaireAnswerPushAPIResponse struct {
    model.CommonResponse
    TmallTmicQuestionnaireAnswerPushResponse
}

type TmallTmicQuestionnaireAnswerPushResponse struct {
    XMLName xml.Name `xml:"tmall_tmic_questionnaire_answer_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误提示语
    
    BizErrorInfo   string `json:"biz_error_info,omitempty" xml:"biz_error_info,omitempty"`

    
    // 错误编码
    
    BizErrorCode   string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`

    
    // 是否调用成功
    
    BizSuccess   bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`

    
}
