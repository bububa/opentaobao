package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
码商查询淘宝码接口 
taobao.eticket.merchant.tbma.get

码商查询淘宝码接口
*/
func TaobaoEticketMerchantTbmaGet(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantTbmaGetRequest, session string) (*eticket.TaobaoEticketMerchantTbmaGetResponse, error) {
    var resp eticket.TaobaoEticketMerchantTbmaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
