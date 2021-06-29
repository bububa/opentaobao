package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
商家批量更新宝贝价格 
taobao.drug.price.batch.update

商家批量更新宝贝价格
*/
func TaobaoDrugPriceBatchUpdate(clt *core.SDKClient, req *alihealth2.TaobaoDrugPriceBatchUpdateRequest, session string) (*alihealth2.TaobaoDrugPriceBatchUpdateAPIResponse, error) {
    var resp alihealth2.TaobaoDrugPriceBatchUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
