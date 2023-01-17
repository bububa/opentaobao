package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest 根据条件删除订阅关系 API请求
// taobao.baichuan.items.unsubscribe.by.condition
//
// 根据条件删除订阅关系
type TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest struct {
	model.Params
	// 删除条件
	_condition *Condition
}

// NewTaobaoBaichuanItemsUnsubscribeByConditionRequest 初始化TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest对象
func NewTaobaoBaichuanItemsUnsubscribeByConditionRequest() *TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest {
	return &TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.items.unsubscribe.by.condition"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCondition is Condition Setter
// 删除条件
func (r *TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest) SetCondition(_condition *Condition) error {
	r._condition = _condition
	r.Set("condition", _condition)
	return nil
}

// GetCondition Condition Getter
func (r TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest) GetCondition() *Condition {
	return r._condition
}
