package wlbimports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlbimports"
)

/* 
物流订单获取 
taobao.wlb.imports.order.get

一般进口物流订单获取
*/
func TaobaoWlbImportsOrderGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsOrderGetRequest, session string) (*wlbimports.TaobaoWlbImportsOrderGetResponse, error) {
    var resp wlbimports.TaobaoWlbImportsOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
