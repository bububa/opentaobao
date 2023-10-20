package alink

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkmessageconfiglistAPIResponse 查询消息开关配置列表 API返回值
// alibaba.alink.message.config.list
//
// 阿里智能获取消息开关配置列表
type AlibabaalinkmessageconfiglistAPIResponse struct {
	model.CommonResponse
	AlibabaalinkmessageconfiglistAPIResponseModel
}

// AlibabaalinkmessageconfiglistAPIResponseModel is 查询消息开关配置列表 成功返回结果
type AlibabaalinkmessageconfiglistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_config_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
