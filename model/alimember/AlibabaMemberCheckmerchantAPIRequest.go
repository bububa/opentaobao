package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验商家身份 API请求
alibaba.member.checkmerchant

校验商家身份
*/
type AlibabaMemberCheckmerchantAPIRequest struct {
    model.Params
    // 混淆后的商家id
    _openMerchantId   string
}

// 初始化AlibabaMemberCheckmerchantAPIRequest对象
func NewAlibabaMemberCheckmerchantRequest() *AlibabaMemberCheckmerchantAPIRequest{
    return &AlibabaMemberCheckmerchantAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMemberCheckmerchantAPIRequest) GetApiMethodName() string {
    return "alibaba.member.checkmerchant"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMemberCheckmerchantAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenMerchantId Setter
// 混淆后的商家id
func (r *AlibabaMemberCheckmerchantAPIRequest) SetOpenMerchantId(_openMerchantId string) error {
    r._openMerchantId = _openMerchantId
    r.Set("open_merchant_id", _openMerchantId)
    return nil
}

// OpenMerchantId Getter
func (r AlibabaMemberCheckmerchantAPIRequest) GetOpenMerchantId() string {
    return r._openMerchantId
}
