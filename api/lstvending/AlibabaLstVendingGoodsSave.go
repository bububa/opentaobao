package lstvending

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstvending"
)

/* 
自动售卖机商品回传 
alibaba.lst.vending.goods.save

零售通自动售卖机商品数据回流。
*/
func AlibabaLstVendingGoodsSave(clt *core.SDKClient, req *lstvending.AlibabaLstVendingGoodsSaveAPIRequest, session string) (*lstvending.AlibabaLstVendingGoodsSaveAPIResponse, error) {
    var resp lstvending.AlibabaLstVendingGoodsSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
