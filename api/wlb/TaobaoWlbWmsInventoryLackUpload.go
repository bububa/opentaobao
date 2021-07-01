package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
缺货通知 
taobao.wlb.wms.inventory.lack.upload

缺货通知
*/
func TaobaoWlbWmsInventoryLackUpload(clt *core.SDKClient, req *wlb.TaobaoWlbWmsInventoryLackUploadAPIRequest, session string) (*wlb.TaobaoWlbWmsInventoryLackUploadAPIResponse, error) {
    var resp wlb.TaobaoWlbWmsInventoryLackUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
