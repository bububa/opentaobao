package omniorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniDealerOdersListAPIRequest) Reset() {
	r._queryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersListAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniDealerOdersListAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoOmniDealerOdersListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniDealerOdersListRequest()
	},
}

// GetTaobaoOmniDealerOdersListRequest 从 sync.Pool 获取 TaobaoOmniDealerOdersListAPIRequest
func GetTaobaoOmniDealerOdersListAPIRequest() *TaobaoOmniDealerOdersListAPIRequest {
	return poolTaobaoOmniDealerOdersListAPIRequest.Get().(*TaobaoOmniDealerOdersListAPIRequest)
}

// ReleaseTaobaoOmniDealerOdersListAPIRequest 将 TaobaoOmniDealerOdersListAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniDealerOdersListAPIRequest(v *TaobaoOmniDealerOdersListAPIRequest) {
	v.Reset()
	poolTaobaoOmniDealerOdersListAPIRequest.Put(v)
}
