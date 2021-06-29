package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三级商户进件修改 APIRequest
tmall.nrt.pay.merchant.stall.signing.modify

三级商户进件修改
*/
type TmallNrtPayMerchantStallSigningModifyRequest struct {
    model.Params

    // 请求参数
    req   *StallSigningReqDto 

}

func NewTmallNrtPayMerchantStallSigningModifyRequest() *TmallNrtPayMerchantStallSigningModifyRequest{
    return &TmallNrtPayMerchantStallSigningModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtPayMerchantStallSigningModifyRequest) GetApiMethodName() string {
    return "tmall.nrt.pay.merchant.stall.signing.modify"
}

func (r TmallNrtPayMerchantStallSigningModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtPayMerchantStallSigningModifyRequest) SetReq(req *StallSigningReqDto) error {
    r.req = req
    r.Set("req", req)
    return nil
}

func (r TmallNrtPayMerchantStallSigningModifyRequest) GetReq() *StallSigningReqDto {
    return r.req
}

