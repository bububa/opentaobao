package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointDecreaseAPIRequest 互动积分扣减接口 API请求
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
type TaobaoJstInteractivePointDecreaseAPIRequest struct {
	model.Params
	// 扣减的积分值
	_amount int64
	// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
	_operateCode string
}

// NewTaobaoJstInteractivePointDecreaseRequest 初始化TaobaoJstInteractivePointDecreaseAPIRequest对象
func NewTaobaoJstInteractivePointDecreaseRequest() *TaobaoJstInteractivePointDecreaseAPIRequest {
	return &TaobaoJstInteractivePointDecreaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.decrease"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAmount is Amount Setter
// 扣减的积分值
func (r *TaobaoJstInteractivePointDecreaseAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetAmount() int64 {
	return r._amount
}

// SetOperateCode is OperateCode Setter
// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
func (r *TaobaoJstInteractivePointDecreaseAPIRequest) SetOperateCode(_operateCode string) error {
	r._operateCode = _operateCode
	r.Set("operate_code", _operateCode)
	return nil
}

// GetOperateCode OperateCode Getter
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetOperateCode() string {
	return r._operateCode
}
