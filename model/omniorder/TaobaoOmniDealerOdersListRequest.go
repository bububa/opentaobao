package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道经销商订单列表 APIRequest
taobao.omni.dealer.oders.list

全渠道经销商订单列表查询
*/
type TaobaoOmniDealerOdersListRequest struct {
    model.Params

    // 参数对象
    queryParam   *QueryOmniOrderRequest 

}

func NewTaobaoOmniDealerOdersListRequest() *TaobaoOmniDealerOdersListRequest{
    return &TaobaoOmniDealerOdersListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniDealerOdersListRequest) GetApiMethodName() string {
    return "taobao.omni.dealer.oders.list"
}

func (r TaobaoOmniDealerOdersListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniDealerOdersListRequest) SetQueryParam(queryParam *QueryOmniOrderRequest) error {
    r.queryParam = queryParam
    r.Set("query_param", queryParam)
    return nil
}

func (r TaobaoOmniDealerOdersListRequest) GetQueryParam() *QueryOmniOrderRequest {
    return r.queryParam
}

