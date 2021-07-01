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
type AlibabaTclsAelophyMerchantIdMixAPIRequest struct {
    model.Params
    // 商家用户id
    _unionUid   string
}

// 初始化AlibabaTclsAelophyMerchantIdMixAPIRequest对象
func NewAlibabaTclsAelophyMerchantIdMixRequest() *AlibabaTclsAelophyMerchantIdMixAPIRequest{
    return &AlibabaTclsAelophyMerchantIdMixAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.id.mix"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnionUid Setter
// 商家用户id
func (r *AlibabaTclsAelophyMerchantIdMixAPIRequest) SetUnionUid(_unionUid string) error {
    r._unionUid = _unionUid
    r.Set("union_uid", _unionUid)
    return nil
}

// UnionUid Getter
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetUnionUid() string {
    return r._unionUid
}
