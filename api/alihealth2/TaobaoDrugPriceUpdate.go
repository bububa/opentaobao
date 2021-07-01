package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
商家更新宝贝价格 
taobao.drug.price.update

商家更新价格
*/
func TaobaoDrugPriceUpdate(clt *core.SDKClient, req *alihealth2.TaobaoDrugPriceUpdateAPIRequest, session string) (*alihealth2.TaobaoDrugPriceUpdateAPIResponse, error) {
    var resp alihealth2.TaobaoDrugPriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
