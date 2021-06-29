package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
不落库商家推送更新酒店rate 
taobao.xhotel.intl.rate.update

商家主动推送不落库商品的酒店信息
*/
func TaobaoXhotelIntlRateUpdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelIntlRateUpdateRequest, session string) (*xhotelonlineorder.TaobaoXhotelIntlRateUpdateAPIResponse, error) {
    var resp xhotelonlineorder.TaobaoXhotelIntlRateUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
