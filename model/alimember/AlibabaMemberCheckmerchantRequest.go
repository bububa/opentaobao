package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验商家身份 APIRequest
alibaba.member.checkmerchant

校验商家身份
*/
type AlibabaMemberCheckmerchantRequest struct {
    model.Params

    // 混淆后的商家id
    openMerchantId   string 

}

func NewAlibabaMemberCheckmerchantRequest() *AlibabaMemberCheckmerchantRequest{
    return &AlibabaMemberCheckmerchantRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMemberCheckmerchantRequest) GetApiMethodName() string {
    return "alibaba.member.checkmerchant"
}

func (r AlibabaMemberCheckmerchantRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMemberCheckmerchantRequest) SetOpenMerchantId(openMerchantId string) error {
    r.openMerchantId = openMerchantId
    r.Set("open_merchant_id", openMerchantId)
    return nil
}

func (r AlibabaMemberCheckmerchantRequest) GetOpenMerchantId() string {
    return r.openMerchantId
}

