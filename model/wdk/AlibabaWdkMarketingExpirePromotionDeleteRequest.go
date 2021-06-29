package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠删除 API请求
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除
*/
type AlibabaWdkMarketingExpirePromotionDeleteRequest struct {
    model.Params
    // 删除短保优惠
    _param0   *ExpirePromotionBo
}

// 初始化AlibabaWdkMarketingExpirePromotionDeleteRequest对象
func NewAlibabaWdkMarketingExpirePromotionDeleteRequest() *AlibabaWdkMarketingExpirePromotionDeleteRequest{
    return &AlibabaWdkMarketingExpirePromotionDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionDeleteRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.expire.promotion.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 删除短保优惠
func (r *AlibabaWdkMarketingExpirePromotionDeleteRequest) SetParam0(_param0 *ExpirePromotionBo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingExpirePromotionDeleteRequest) GetParam0() *ExpirePromotionBo {
    return r._param0
}
