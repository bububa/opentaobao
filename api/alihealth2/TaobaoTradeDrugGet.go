package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询商家未确认订单列表 
taobao.trade.drug.get

可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
*/
func TaobaoTradeDrugGet(clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugGetRequest, session string) (*alihealth2.TaobaoTradeDrugGetAPIResponse, error) {
    var resp alihealth2.TaobaoTradeDrugGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
