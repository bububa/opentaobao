package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
卖家取消店铺vip的优惠 
taobao.crm.shopvip.cancel

此接口用于取消VIP优惠
*/
func TaobaoCrmShopvipCancel(clt *core.SDKClient, req *crm.TaobaoCrmShopvipCancelRequest, session string) (*crm.TaobaoCrmShopvipCancelResponse, error) {
    var resp crm.TaobaoCrmShopvipCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
