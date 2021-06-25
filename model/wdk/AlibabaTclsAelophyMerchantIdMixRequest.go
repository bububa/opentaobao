package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家用户id混淆 APIRequest
alibaba.tcls.aelophy.merchant.id.mix

商家用户id混淆
*/
type AlibabaTclsAelophyMerchantIdMixRequest struct {
    model.Params

    // 商家用户id
    unionUid   string 

}

func NewAlibabaTclsAelophyMerchantIdMixRequest() *AlibabaTclsAelophyMerchantIdMixRequest{
    return &AlibabaTclsAelophyMerchantIdMixRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantIdMixRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.id.mix"
}

func (r AlibabaTclsAelophyMerchantIdMixRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantIdMixRequest) SetUnionUid(unionUid string) error {
    r.unionUid = unionUid
    r.Set("union_uid", unionUid)
    return nil
}

func (r AlibabaTclsAelophyMerchantIdMixRequest) GetUnionUid() string {
    return r.unionUid
}

