package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniusrefundscheckAPIRequest 退款信息审核 API请求
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
type TaobaordcaligeniusrefundscheckAPIRequest struct {
	model.Params
	// 入参
	_param *RefundCheckDto
}

// NewTaobaordcaligeniusrefundscheckRequest 初始化TaobaordcaligeniusrefundscheckAPIRequest对象
func NewTaobaordcaligeniusrefundscheckRequest() *TaobaordcaligeniusrefundscheckAPIRequest {
	return &TaobaordcaligeniusrefundscheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniusrefundscheckAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.refunds.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniusrefundscheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniusrefundscheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaordcaligeniusrefundscheckAPIRequest) SetParam(_param *RefundCheckDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaordcaligeniusrefundscheckAPIRequest) GetParam() *RefundCheckDto {
	return r._param
}
