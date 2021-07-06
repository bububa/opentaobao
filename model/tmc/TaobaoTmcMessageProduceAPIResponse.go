package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessageProduceAPIResponse 发布单条消息 API返回值
// taobao.tmc.message.produce
//
// 发布单条消息
type TaobaoTmcMessageProduceAPIResponse struct {
	model.CommonResponse
	TaobaoTmcMessageProduceAPIResponseModel
}

// TaobaoTmcMessageProduceAPIResponseModel is 发布单条消息 成功返回结果
type TaobaoTmcMessageProduceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_message_produce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息ID
	MsgIds []string `json:"msg_ids,omitempty" xml:"msg_ids>string,omitempty"`
	// 投递目标数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
