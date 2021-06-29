package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三级商户进件修改 API请求
tmall.nrt.pay.merchant.stall.signing.modify

三级商户进件修改
*/
type TmallNrtPayMerchantStallSigningModifyRequest struct {
    model.Params
    // 请求参数
    _req   *StallSigningReqDto
}

// 初始化TmallNrtPayMerchantStallSigningModifyRequest对象
func NewTmallNrtPayMerchantStallSigningModifyRequest() *TmallNrtPayMerchantStallSigningModifyRequest{
    return &TmallNrtPayMerchantStallSigningModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtPayMerchantStallSigningModifyRequest) GetApiMethodName() string {
    return "tmall.nrt.pay.merchant.stall.signing.modify"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtPayMerchantStallSigningModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Req Setter
// 请求参数
func (r *TmallNrtPayMerchantStallSigningModifyRequest) SetReq(_req *StallSigningReqDto) error {
    r._req = _req
    r.Set("req", _req)
    return nil
}

// Req Getter
func (r TmallNrtPayMerchantStallSigningModifyRequest) GetReq() *StallSigningReqDto {
    return r._req
}
