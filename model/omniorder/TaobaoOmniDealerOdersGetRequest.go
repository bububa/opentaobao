package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔全渠道经销商订单的详细信息 APIRequest
taobao.omni.dealer.oders.get

全渠道经销商获取单笔订单的详细信息
*/
type TaobaoOmniDealerOdersGetRequest struct {
    model.Params

    // 主订单ID
    mainOrderId   int64 

}

func NewTaobaoOmniDealerOdersGetRequest() *TaobaoOmniDealerOdersGetRequest{
    return &TaobaoOmniDealerOdersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniDealerOdersGetRequest) GetApiMethodName() string {
    return "taobao.omni.dealer.oders.get"
}

func (r TaobaoOmniDealerOdersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniDealerOdersGetRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoOmniDealerOdersGetRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

