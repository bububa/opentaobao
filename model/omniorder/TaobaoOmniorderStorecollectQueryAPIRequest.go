package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorecollectqueryAPIRequest 全渠道门店自提根据核销码查询订单 API请求
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
type TaobaoomniorderstorecollectqueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// NewTaobaoomniorderstorecollectqueryRequest 初始化TaobaoomniorderstorecollectqueryAPIRequest对象
func NewTaobaoomniorderstorecollectqueryRequest() *TaobaoomniorderstorecollectqueryAPIRequest {
	return &TaobaoomniorderstorecollectqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstorecollectqueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.storecollect.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstorecollectqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstorecollectqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoomniorderstorecollectqueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoomniorderstorecollectqueryAPIRequest) GetCode() string {
	return r._code
}
