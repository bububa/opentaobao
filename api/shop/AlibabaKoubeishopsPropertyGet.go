package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
口碑店铺列表推荐 
alibaba.koubeishops.property.get

推荐用户附近的美食门店
*/
func AlibabaKoubeishopsPropertyGet(clt *core.SDKClient, req *shop.AlibabaKoubeishopsPropertyGetAPIRequest, session string) (*shop.AlibabaKoubeishopsPropertyGetAPIResponse, error) {
    var resp shop.AlibabaKoubeishopsPropertyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
