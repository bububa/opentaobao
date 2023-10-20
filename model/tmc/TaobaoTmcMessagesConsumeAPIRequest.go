package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesConsumeAPIRequest 消费多条消息 API请求
// taobao.tmc.messages.consume
//
// 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaoTmcMessagesConsumeAPIRequest struct {
	model.Params
	// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
	_groupName string
	// 每次批量消费消息的条数，最小值：10；最大值：200
	_quantity int64
}

// NewTaobaoTmcMessagesConsumeRequest 初始化TaobaoTmcMessagesConsumeAPIRequest对象
func NewTaobaoTmcMessagesConsumeRequest() *TaobaoTmcMessagesConsumeAPIRequest {
	return &TaobaoTmcMessagesConsumeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcMessagesConsumeAPIRequest) Reset() {
	r._groupName = ""
	r._quantity = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcMessagesConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
func (r *TaobaoTmcMessagesConsumeAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaoTmcMessagesConsumeAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetQuantity is Quantity Setter
// 每次批量消费消息的条数，最小值：10；最大值：200
func (r *TaobaoTmcMessagesConsumeAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoTmcMessagesConsumeAPIRequest) GetQuantity() int64 {
	return r._quantity
}

var poolTaobaoTmcMessagesConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcMessagesConsumeRequest()
	},
}

// GetTaobaoTmcMessagesConsumeRequest 从 sync.Pool 获取 TaobaoTmcMessagesConsumeAPIRequest
func GetTaobaoTmcMessagesConsumeAPIRequest() *TaobaoTmcMessagesConsumeAPIRequest {
	return poolTaobaoTmcMessagesConsumeAPIRequest.Get().(*TaobaoTmcMessagesConsumeAPIRequest)
}

// ReleaseTaobaoTmcMessagesConsumeAPIRequest 将 TaobaoTmcMessagesConsumeAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcMessagesConsumeAPIRequest(v *TaobaoTmcMessagesConsumeAPIRequest) {
	v.Reset()
	poolTaobaoTmcMessagesConsumeAPIRequest.Put(v)
}
