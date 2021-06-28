package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
电子凭证核销接口 
taobao.eticket.merchant.ma.consume

电子凭证核销接口
*/
func TaobaoEticketMerchantMaConsume(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaConsumeRequest, session string) (*eticket.TaobaoEticketMerchantMaConsumeAPIResponse, error) {
    var resp eticket.TaobaoEticketMerchantMaConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
