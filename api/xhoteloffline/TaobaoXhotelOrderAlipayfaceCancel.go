package xhoteloffline

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhoteloffline"
)

/* 
线下信用住订单取消 
taobao.xhotel.order.alipayface.cancel

提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
*/
func TaobaoXhotelOrderAlipayfaceCancel(clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderAlipayfaceCancelAPIRequest, session string) (*xhoteloffline.TaobaoXhotelOrderAlipayfaceCancelAPIResponse, error) {
    var resp xhoteloffline.TaobaoXhotelOrderAlipayfaceCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
