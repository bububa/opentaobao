package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStorecollectQueryAPIRequest 全渠道门店自提根据核销码查询订单 API请求
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// NewTaobaoOmniorderStorecollectQueryRequest 初始化TaobaoOmniorderStorecollectQueryAPIRequest对象
func NewTaobaoOmniorderStorecollectQueryRequest() *TaobaoOmniorderStorecollectQueryAPIRequest {
	return &TaobaoOmniorderStorecollectQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.storecollect.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoOmniorderStorecollectQueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetCode() string {
	return r._code
}
