package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractivepointdecreaseAPIRequest 互动积分扣减接口 API请求
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
type TaobaojstinteractivepointdecreaseAPIRequest struct {
	model.Params
	// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
	_operateCode string
	// 扣减的积分值
	_amount int64
}

// NewTaobaojstinteractivepointdecreaseRequest 初始化TaobaojstinteractivepointdecreaseAPIRequest对象
func NewTaobaojstinteractivepointdecreaseRequest() *TaobaojstinteractivepointdecreaseAPIRequest {
	return &TaobaojstinteractivepointdecreaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractivepointdecreaseAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.decrease"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractivepointdecreaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractivepointdecreaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperateCode is OperateCode Setter
// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
func (r *TaobaojstinteractivepointdecreaseAPIRequest) SetOperateCode(_operateCode string) error {
	r._operateCode = _operateCode
	r.Set("operate_code", _operateCode)
	return nil
}

// GetOperateCode OperateCode Getter
func (r TaobaojstinteractivepointdecreaseAPIRequest) GetOperateCode() string {
	return r._operateCode
}

// SetAmount is Amount Setter
// 扣减的积分值
func (r *TaobaojstinteractivepointdecreaseAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r TaobaojstinteractivepointdecreaseAPIRequest) GetAmount() int64 {
	return r._amount
}
