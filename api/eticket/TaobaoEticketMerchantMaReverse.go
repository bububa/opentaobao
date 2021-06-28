package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
电子凭证冲正接口 
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口
*/
func TaobaoEticketMerchantMaReverse(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaReverseRequest, session string) (*eticket.TaobaoEticketMerchantMaReverseAPIResponse, error) {
    var resp eticket.TaobaoEticketMerchantMaReverseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
