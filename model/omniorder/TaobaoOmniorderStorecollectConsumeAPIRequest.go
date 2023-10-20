package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorecollectconsumeAPIRequest 全渠道门店自提核销订单 API请求
// taobao.omniorder.storecollect.consume
//
// 全渠道门店自提核销订单
type TaobaoomniorderstorecollectconsumeAPIRequest struct {
	model.Params
	// 核销码
	_code string
	// 核销操作人信息
	_operator string
	// 淘宝主订单ID
	_mainOrderId int64
}

// NewTaobaoomniorderstorecollectconsumeRequest 初始化TaobaoomniorderstorecollectconsumeAPIRequest对象
func NewTaobaoomniorderstorecollectconsumeRequest() *TaobaoomniorderstorecollectconsumeAPIRequest {
	return &TaobaoomniorderstorecollectconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstorecollectconsumeAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.storecollect.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstorecollectconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstorecollectconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoomniorderstorecollectconsumeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoomniorderstorecollectconsumeAPIRequest) GetCode() string {
	return r._code
}

// SetOperator is Operator Setter
// 核销操作人信息
func (r *TaobaoomniorderstorecollectconsumeAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoomniorderstorecollectconsumeAPIRequest) GetOperator() string {
	return r._operator
}

// SetMainOrderId is MainOrderId Setter
// 淘宝主订单ID
func (r *TaobaoomniorderstorecollectconsumeAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoomniorderstorecollectconsumeAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
