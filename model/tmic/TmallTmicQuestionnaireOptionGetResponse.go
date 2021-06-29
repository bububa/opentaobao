package tmic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单题选项 API返回值 
tmall.tmic.questionnaire.option.get

根据具体题号，获取当前题目的选项列表
*/
type TmallTmicQuestionnaireOptionGetAPIResponse struct {
    model.CommonResponse
    TmallTmicQuestionnaireOptionGetResponse
}

// 获取单题选项 成功返回结果
type TmallTmicQuestionnaireOptionGetResponse struct {
    XMLName xml.Name `xml:"tmall_tmic_questionnaire_option_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *OpenOptionResult `json:"result,omitempty" xml:"result,omitempty"`
}
