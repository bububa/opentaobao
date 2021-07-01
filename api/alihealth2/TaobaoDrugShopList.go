package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询卖家外卖店列表 
taobao.drug.shop.list

查询卖家外卖店列表
*/
func TaobaoDrugShopList(clt *core.SDKClient, req *alihealth2.TaobaoDrugShopListAPIRequest, session string) (*alihealth2.TaobaoDrugShopListAPIResponse, error) {
    var resp alihealth2.TaobaoDrugShopListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
