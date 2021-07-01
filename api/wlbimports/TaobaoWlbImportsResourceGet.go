package wlbimports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlbimports"
)

/* 
获取所有服务列表 
taobao.wlb.imports.resource.get

一般进口TOP接口，获取所有服务列表。
*/
func TaobaoWlbImportsResourceGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsResourceGetAPIRequest, session string) (*wlbimports.TaobaoWlbImportsResourceGetAPIResponse, error) {
    var resp wlbimports.TaobaoWlbImportsResourceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
