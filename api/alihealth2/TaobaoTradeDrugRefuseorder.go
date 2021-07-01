package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
阿里健康020拒单 
taobao.trade.drug.refuseorder

阿里健康020拒单
*/
func TaobaoTradeDrugRefuseorder(clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugRefuseorderAPIRequest, session string) (*alihealth2.TaobaoTradeDrugRefuseorderAPIResponse, error) {
    var resp alihealth2.TaobaoTradeDrugRefuseorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
