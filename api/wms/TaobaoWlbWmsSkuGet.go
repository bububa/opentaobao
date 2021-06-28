package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
商品信息查询 
taobao.wlb.wms.sku.get

商品信息查询
*/
func TaobaoWlbWmsSkuGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuGetRequest, session string) (*wms.TaobaoWlbWmsSkuGetAPIResponse, error) {
    var resp wms.TaobaoWlbWmsSkuGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
