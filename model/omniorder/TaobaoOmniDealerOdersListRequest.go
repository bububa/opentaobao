package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道经销商订单列表 API请求
taobao.omni.dealer.oders.list

全渠道经销商订单列表查询
*/
type TaobaoOmniDealerOdersListRequest struct {
    model.Params
    // 参数对象
    _queryParam   *QueryOmniOrderRequest
}

// 初始化TaobaoOmniDealerOdersListRequest对象
func NewTaobaoOmniDealerOdersListRequest() *TaobaoOmniDealerOdersListRequest{
    return &TaobaoOmniDealerOdersListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersListRequest) GetApiMethodName() string {
    return "taobao.omni.dealer.oders.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryParam Setter
// 参数对象
func (r *TaobaoOmniDealerOdersListRequest) SetQueryParam(_queryParam *QueryOmniOrderRequest) error {
    r._queryParam = _queryParam
    r.Set("query_param", _queryParam)
    return nil
}

// QueryParam Getter
func (r TaobaoOmniDealerOdersListRequest) GetQueryParam() *QueryOmniOrderRequest {
    return r._queryParam
}
