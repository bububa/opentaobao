package tmc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcMessagesProduceAPIRequest) Reset() {
	r._messages = r._messages[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesProduceAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.produce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesProduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcMessagesProduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessages is Messages Setter
// tmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{&#34;content&#34;: &#34;{\&#34;tid\&#34;:1234554321,\&#34;status\&#34;:\&#34;X_LOGISTICS_PRINTED\&#34;,\&#34;action_time\&#34;:\&#34;2014-08-08 18:24:00\&#34;,\&#34;seller_nick\&#34;: \&#34;向阳aa\&#34;,\&#34;operator\&#34;:\&#34;小张\&#34;}&#34;,&#34;topic&#34;: &#34;taobao_jds_TradeTrace&#34;},{&#34;content&#34;: &#34;{\&#34;tid\&#34;:1234554321,\&#34;status\&#34;:\&#34;X_LOGISTICS_PRINTED\&#34;,\&#34;action_time\&#34;:\&#34;2014-08-08 18:24:00\&#34;,\&#34;seller_nick\&#34;: \&#34;向阳aa\&#34;,\&#34;operator\&#34;:\&#34;小张\&#34;}&#34;,&#34;topic&#34;: &#34;taobao_jds_TradeTrace&#34;}]
func (r *TaobaoTmcMessagesProduceAPIRequest) SetMessages(_messages []TmcPublishMessage) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// GetMessages Messages Getter
func (r TaobaoTmcMessagesProduceAPIRequest) GetMessages() []TmcPublishMessage {
	return r._messages
}

var poolTaobaoTmcMessagesProduceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcMessagesProduceRequest()
	},
}

// GetTaobaoTmcMessagesProduceRequest 从 sync.Pool 获取 TaobaoTmcMessagesProduceAPIRequest
func GetTaobaoTmcMessagesProduceAPIRequest() *TaobaoTmcMessagesProduceAPIRequest {
	return poolTaobaoTmcMessagesProduceAPIRequest.Get().(*TaobaoTmcMessagesProduceAPIRequest)
}

// ReleaseTaobaoTmcMessagesProduceAPIRequest 将 TaobaoTmcMessagesProduceAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcMessagesProduceAPIRequest(v *TaobaoTmcMessagesProduceAPIRequest) {
	v.Reset()
	poolTaobaoTmcMessagesProduceAPIRequest.Put(v)
}
