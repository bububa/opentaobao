package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务质量反馈编码列表 API返回值 
cainiao.nbadd.appointdeliver.feedbackcodes

服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
*/
type CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse struct {
    model.CommonResponse
    CainiaoNbaddAppointdeliverFeedbackcodesResponse
}

// 服务质量反馈编码列表 成功返回结果
type CainiaoNbaddAppointdeliverFeedbackcodesResponse struct {
    XMLName xml.Name `xml:"cainiao_nbadd_appointdeliver_feedbackcodes_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误描述
    ResultDesc   string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
    // 错误编码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 接口调用是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 返回的具体数据
    ResultList   []FeedbackCodeDTO `json:"result_list,omitempty" xml:"result_list>feedback_code_dto,omitempty"`
}
