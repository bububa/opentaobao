package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserTopicsGetAPIResponse 获取用户开通的topic列表 API返回值
// taobao.tmc.user.topics.get
//
// 获取用户开通的topic列表，授权消息使用
type TaobaoTmcUserTopicsGetAPIResponse struct {
	model.CommonResponse
	TaobaoTmcUserTopicsGetAPIResponseModel
}

// TaobaoTmcUserTopicsGetAPIResponseModel is 获取用户开通的topic列表 成功返回结果
type TaobaoTmcUserTopicsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_user_topics_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// topic列表
	Topics []string `json:"topics,omitempty" xml:"topics>string,omitempty"`
	// 错误信息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
