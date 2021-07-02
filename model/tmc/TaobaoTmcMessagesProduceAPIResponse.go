package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesProduceAPIResponse 批量发送消息 API返回值
// taobao.tmc.messages.produce
//
// 批量发送消息
type TaobaoTmcMessagesProduceAPIResponse struct {
	model.CommonResponse
	TaobaoTmcMessagesProduceAPIResponseModel
}

// TaobaoTmcMessagesProduceAPIResponseModel is 批量发送消息 成功返回结果
type TaobaoTmcMessagesProduceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_messages_produce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否全部成功
	IsAllSuccess bool `json:"is_all_success,omitempty" xml:"is_all_success,omitempty"`
	// 发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
	Results []TmcProduceResult `json:"results,omitempty" xml:"results>tmc_produce_result,omitempty"`
}
