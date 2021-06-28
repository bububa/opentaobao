package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
商品信息的更新 
taobao.wlb.wms.sku.update

商品信息的更新
*/
func TaobaoWlbWmsSkuUpdate(clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuUpdateRequest, session string) (*wms.TaobaoWlbWmsSkuUpdateAPIResponse, error) {
    var resp wms.TaobaoWlbWmsSkuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
