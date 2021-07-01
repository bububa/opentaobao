package xhoteloffline

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhoteloffline"
)

/* 
信用住支付创建接口 
taobao.xhotel.order.alipayface.create

用于创建一笔信用住支付，主要应用场景是线下信用住
*/
func TaobaoXhotelOrderAlipayfaceCreate(clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderAlipayfaceCreateAPIRequest, session string) (*xhoteloffline.TaobaoXhotelOrderAlipayfaceCreateAPIResponse, error) {
    var resp xhoteloffline.TaobaoXhotelOrderAlipayfaceCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
