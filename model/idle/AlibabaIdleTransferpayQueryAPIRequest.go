package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTransferpayQueryAPIRequest 闲鱼转账结果查询 API请求
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
type AlibabaIdleTransferpayQueryAPIRequest struct {
	model.Params
	// 入参
	_param *PayQueryRequest
}

// NewAlibabaIdleTransferpayQueryRequest 初始化AlibabaIdleTransferpayQueryAPIRequest对象
func NewAlibabaIdleTransferpayQueryRequest() *AlibabaIdleTransferpayQueryAPIRequest {
	return &AlibabaIdleTransferpayQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTransferpayQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.transferpay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTransferpayQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaIdleTransferpayQueryAPIRequest) SetParam(_param *PayQueryRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleTransferpayQueryAPIRequest) GetParam() *PayQueryRequest {
	return r._param
}
