package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠创建 API请求
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入
*/
type AlibabaWdkMarketingExpirePromotionCreateRequest struct {
    model.Params
    // 创建短保优惠
    param0   *ExpirePromotionBo
}

// 初始化AlibabaWdkMarketingExpirePromotionCreateRequest对象
func NewAlibabaWdkMarketingExpirePromotionCreateRequest() *AlibabaWdkMarketingExpirePromotionCreateRequest{
    return &AlibabaWdkMarketingExpirePromotionCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.expire.promotion.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 创建短保优惠
func (r *AlibabaWdkMarketingExpirePromotionCreateRequest) SetParam0(param0 *ExpirePromotionBo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingExpirePromotionCreateRequest) GetParam0() *ExpirePromotionBo {
    return r.param0
}
