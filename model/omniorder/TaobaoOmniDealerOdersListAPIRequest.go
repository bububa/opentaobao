package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomnidealeroderslistAPIRequest 全渠道经销商订单列表 API请求
// taobao.omni.dealer.oders.list
//
// 全渠道经销商订单列表查询
type TaobaoomnidealeroderslistAPIRequest struct {
	model.Params
	// 参数对象
	_queryParam *QueryOmniOrderRequest
}

// NewTaobaoomnidealeroderslistRequest 初始化TaobaoomnidealeroderslistAPIRequest对象
func NewTaobaoomnidealeroderslistRequest() *TaobaoomnidealeroderslistAPIRequest {
	return &TaobaoomnidealeroderslistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomnidealeroderslistAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomnidealeroderslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomnidealeroderslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryParam is QueryParam Setter
// 参数对象
func (r *TaobaoomnidealeroderslistAPIRequest) SetQueryParam(_queryParam *QueryOmniOrderRequest) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r TaobaoomnidealeroderslistAPIRequest) GetQueryParam() *QueryOmniOrderRequest {
	return r._queryParam
}
