package wlbimports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlbimports"
)

/* 
获取运单信息 
taobao.wlb.imports.waybill.get

一般进口商家，获取订单的电子面单链接地址
*/
func TaobaoWlbImportsWaybillGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsWaybillGetRequest, session string) (*wlbimports.TaobaoWlbImportsWaybillGetResponse, error) {
    var resp wlbimports.TaobaoWlbImportsWaybillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
