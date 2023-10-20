package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionorderreceiptinfogetAPIRequest Get Order Receipt Info API请求
// aliexpress.solution.order.receiptinfo.get
//
// Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
type AliexpresssolutionorderreceiptinfogetAPIRequest struct {
	model.Params
	// query param
	_param1 *SingleOrderQuery
}

// NewAliexpresssolutionorderreceiptinfogetRequest 初始化AliexpresssolutionorderreceiptinfogetAPIRequest对象
func NewAliexpresssolutionorderreceiptinfogetRequest() *AliexpresssolutionorderreceiptinfogetAPIRequest {
	return &AliexpresssolutionorderreceiptinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionorderreceiptinfogetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.order.receiptinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionorderreceiptinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionorderreceiptinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// query param
func (r *AliexpresssolutionorderreceiptinfogetAPIRequest) SetParam1(_param1 *SingleOrderQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresssolutionorderreceiptinfogetAPIRequest) GetParam1() *SingleOrderQuery {
	return r._param1
}
