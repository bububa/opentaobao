package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
更新店铺基本信息 
taobao.shop.update

目前只支持标题、公告和描述的更新
*/
func TaobaoShopUpdate(clt *core.SDKClient, req *shop.TaobaoShopUpdateRequest, session string) (*shop.TaobaoShopUpdateAPIResponse, error) {
    var resp shop.TaobaoShopUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
