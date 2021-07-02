package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesProduceAPIRequest 批量发送消息 API请求
// taobao.tmc.messages.produce
//
// 批量发送消息
type TaobaoTmcMessagesProduceAPIRequest struct {
	model.Params
	// tmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}]
	_messages []TmcPublishMessage
}

// NewTaobaoTmcMessagesProduceRequest 初始化TaobaoTmcMessagesProduceAPIRequest对象
func NewTaobaoTmcMessagesProduceRequest() *TaobaoTmcMessagesProduceAPIRequest {
	return &TaobaoTmcMessagesProduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesProduceAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.produce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesProduceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Messages Setter
// tmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}]
func (r *TaobaoTmcMessagesProduceAPIRequest) SetMessages(_messages []TmcPublishMessage) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// Get Messages Getter
func (r TaobaoTmcMessagesProduceAPIRequest) GetMessages() []TmcPublishMessage {
	return r._messages
}
