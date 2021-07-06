package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointIncreaseAPIRequest 互动积分发放接口 API请求
// taobao.jst.interactive.point.increase
//
// 向用户发放互动积分
type TaobaoJstInteractivePointIncreaseAPIRequest struct {
	model.Params
	// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
	_operateCode string
	// 发放的积分值
	_amount int64
}

// NewTaobaoJstInteractivePointIncreaseRequest 初始化TaobaoJstInteractivePointIncreaseAPIRequest对象
func NewTaobaoJstInteractivePointIncreaseRequest() *TaobaoJstInteractivePointIncreaseAPIRequest {
	return &TaobaoJstInteractivePointIncreaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointIncreaseAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.increase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointIncreaseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOperateCode is OperateCode Setter
// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
func (r *TaobaoJstInteractivePointIncreaseAPIRequest) SetOperateCode(_operateCode string) error {
	r._operateCode = _operateCode
	r.Set("operate_code", _operateCode)
	return nil
}

// GetOperateCode OperateCode Getter
func (r TaobaoJstInteractivePointIncreaseAPIRequest) GetOperateCode() string {
	return r._operateCode
}

// SetAmount is Amount Setter
// 发放的积分值
func (r *TaobaoJstInteractivePointIncreaseAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r TaobaoJstInteractivePointIncreaseAPIRequest) GetAmount() int64 {
	return r._amount
}
