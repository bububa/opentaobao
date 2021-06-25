package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家取消店铺vip的优惠 APIRequest
taobao.crm.shopvip.cancel

此接口用于取消VIP优惠
*/
type TaobaoCrmShopvipCancelRequest struct {
    model.Params

}

func NewTaobaoCrmShopvipCancelRequest() *TaobaoCrmShopvipCancelRequest{
    return &TaobaoCrmShopvipCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmShopvipCancelRequest) GetApiMethodName() string {
    return "taobao.crm.shopvip.cancel"
}

func (r TaobaoCrmShopvipCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


