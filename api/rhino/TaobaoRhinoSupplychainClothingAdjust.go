package rhino

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/rhino"
)

/* 
同步成衣仓盘点数据 
taobao.rhino.supplychain.clothing.adjust

同步成衣仓盘点数据
*/
func TaobaoRhinoSupplychainClothingAdjust(clt *core.SDKClient, req *rhino.TaobaoRhinoSupplychainClothingAdjustRequest, session string) (*rhino.TaobaoRhinoSupplychainClothingAdjustAPIResponse, error) {
    var resp rhino.TaobaoRhinoSupplychainClothingAdjustAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}