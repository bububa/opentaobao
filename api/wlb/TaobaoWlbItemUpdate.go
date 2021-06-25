package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
物流宝商品修改 
taobao.wlb.item.update

修改物流宝商品信息
*/
func TaobaoWlbItemUpdate(clt *core.SDKClient, req *wlb.TaobaoWlbItemUpdateRequest, session string) (*wlb.TaobaoWlbItemUpdateResponse, error) {
    var resp wlb.TaobaoWlbItemUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
