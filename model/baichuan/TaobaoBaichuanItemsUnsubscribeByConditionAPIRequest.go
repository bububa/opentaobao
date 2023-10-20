package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsunsubscribebyconditionAPIRequest 根据条件删除订阅关系 API请求
// taobao.baichuan.items.unsubscribe.by.condition
//
// 根据条件删除订阅关系
type TaobaobaichuanitemsunsubscribebyconditionAPIRequest struct {
	model.Params
	// 删除条件
	_condition *Condition
}

// NewTaobaobaichuanitemsunsubscribebyconditionRequest 初始化TaobaobaichuanitemsunsubscribebyconditionAPIRequest对象
func NewTaobaobaichuanitemsunsubscribebyconditionRequest() *TaobaobaichuanitemsunsubscribebyconditionAPIRequest {
	return &TaobaobaichuanitemsunsubscribebyconditionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanitemsunsubscribebyconditionAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.items.unsubscribe.by.condition"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanitemsunsubscribebyconditionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanitemsunsubscribebyconditionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCondition is Condition Setter
// 删除条件
func (r *TaobaobaichuanitemsunsubscribebyconditionAPIRequest) SetCondition(_condition *Condition) error {
	r._condition = _condition
	r.Set("condition", _condition)
	return nil
}

// GetCondition Condition Getter
func (r TaobaobaichuanitemsunsubscribebyconditionAPIRequest) GetCondition() *Condition {
	return r._condition
}
