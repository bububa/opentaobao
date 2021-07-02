package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniDealerOdersListAPIRequest 全渠道经销商订单列表 API请求
// taobao.omni.dealer.oders.list
//
// 全渠道经销商订单列表查询
type TaobaoOmniDealerOdersListAPIRequest struct {
	model.Params
	// 参数对象
	_queryParam *QueryOmniOrderRequest
}

// NewTaobaoOmniDealerOdersListRequest 初始化TaobaoOmniDealerOdersListAPIRequest对象
func NewTaobaoOmniDealerOdersListRequest() *TaobaoOmniDealerOdersListAPIRequest {
	return &TaobaoOmniDealerOdersListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersListAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryParam is QueryParam Setter
// 参数对象
func (r *TaobaoOmniDealerOdersListAPIRequest) SetQueryParam(_queryParam *QueryOmniOrderRequest) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r TaobaoOmniDealerOdersListAPIRequest) GetQueryParam() *QueryOmniOrderRequest {
	return r._queryParam
}
