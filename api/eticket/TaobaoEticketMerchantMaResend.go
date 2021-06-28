package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
电子凭证重发回调接口 
taobao.eticket.merchant.ma.resend

码商重发电子凭证回调接口
*/
func TaobaoEticketMerchantMaResend(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaResendRequest, session string) (*eticket.TaobaoEticketMerchantMaResendAPIResponse, error) {
    var resp eticket.TaobaoEticketMerchantMaResendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
