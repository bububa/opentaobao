package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcmessagesconsumeAPIRequest 消费多条消息 API请求
// taobao.tmc.messages.consume
//
// 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaotmcmessagesconsumeAPIRequest struct {
	model.Params
	// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
	_groupName string
	// 每次批量消费消息的条数，最小值：10；最大值：200
	_quantity int64
}

// NewTaobaotmcmessagesconsumeRequest 初始化TaobaotmcmessagesconsumeAPIRequest对象
func NewTaobaotmcmessagesconsumeRequest() *TaobaotmcmessagesconsumeAPIRequest {
	return &TaobaotmcmessagesconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcmessagesconsumeAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcmessagesconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcmessagesconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
func (r *TaobaotmcmessagesconsumeAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaotmcmessagesconsumeAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetQuantity is Quantity Setter
// 每次批量消费消息的条数，最小值：10；最大值：200
func (r *TaobaotmcmessagesconsumeAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaotmcmessagesconsumeAPIRequest) GetQuantity() int64 {
	return r._quantity
}
