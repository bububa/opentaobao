package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家用户id混淆 API请求
alibaba.tcls.aelophy.merchant.id.mix

商家用户id混淆
*/
type AlibabaTclsAelophyMerchantIdMixRequest struct {
    model.Params
    // 商家用户id
    _unionUid   string
}

// 初始化AlibabaTclsAelophyMerchantIdMixRequest对象
func NewAlibabaTclsAelophyMerchantIdMixRequest() *AlibabaTclsAelophyMerchantIdMixRequest{
    return &AlibabaTclsAelophyMerchantIdMixRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantIdMixRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.id.mix"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantIdMixRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnionUid Setter
// 商家用户id
func (r *AlibabaTclsAelophyMerchantIdMixRequest) SetUnionUid(_unionUid string) error {
    r._unionUid = _unionUid
    r.Set("union_uid", _unionUid)
    return nil
}

// UnionUid Getter
func (r AlibabaTclsAelophyMerchantIdMixRequest) GetUnionUid() string {
    return r._unionUid
}
