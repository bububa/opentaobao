package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointDecreaseAPIRequest 互动积分扣减接口 API请求
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
type TaobaoJstInteractivePointDecreaseAPIRequest struct {
	model.Params
	// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
	_operateCode string
	// 扣减的积分值
	_amount int64
}

// NewTaobaoJstInteractivePointDecreaseRequest 初始化TaobaoJstInteractivePointDecreaseAPIRequest对象
func NewTaobaoJstInteractivePointDecreaseRequest() *TaobaoJstInteractivePointDecreaseAPIRequest {
	return &TaobaoJstInteractivePointDecreaseAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractivePointDecreaseAPIRequest) Reset() {
	r._operateCode = ""
	r._amount = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.decrease"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractivePointDecreaseAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoJstInteractivePointDecreaseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractivePointDecreaseRequest()
	},
}

// GetTaobaoJstInteractivePointDecreaseRequest 从 sync.Pool 获取 TaobaoJstInteractivePointDecreaseAPIRequest
func GetTaobaoJstInteractivePointDecreaseAPIRequest() *TaobaoJstInteractivePointDecreaseAPIRequest {
	return poolTaobaoJstInteractivePointDecreaseAPIRequest.Get().(*TaobaoJstInteractivePointDecreaseAPIRequest)
}

// ReleaseTaobaoJstInteractivePointDecreaseAPIRequest 将 TaobaoJstInteractivePointDecreaseAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractivePointDecreaseAPIRequest(v *TaobaoJstInteractivePointDecreaseAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractivePointDecreaseAPIRequest.Put(v)
}
