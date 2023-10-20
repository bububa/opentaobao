package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletransferpayqueryAPIRequest 闲鱼转账结果查询 API请求
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
type AlibabaidletransferpayqueryAPIRequest struct {
	model.Params
	// 入参
	_param *PayQueryRequest
}

// NewAlibabaidletransferpayqueryRequest 初始化AlibabaidletransferpayqueryAPIRequest对象
func NewAlibabaidletransferpayqueryRequest() *AlibabaidletransferpayqueryAPIRequest {
	return &AlibabaidletransferpayqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletransferpayqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.transferpay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletransferpayqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletransferpayqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaidletransferpayqueryAPIRequest) SetParam(_param *PayQueryRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaidletransferpayqueryAPIRequest) GetParam() *PayQueryRequest {
	return r._param
}
