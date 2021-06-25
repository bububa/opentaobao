package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
商品同步 
taobao.wlb.wms.sku.create

商品同步
*/
func TaobaoWlbWmsSkuCreate(clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuCreateRequest, session string) (*wms.TaobaoWlbWmsSkuCreateResponse, error) {
    var resp wms.TaobaoWlbWmsSkuCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
