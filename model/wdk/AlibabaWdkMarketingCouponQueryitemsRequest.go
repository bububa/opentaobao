package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动下的商品 API请求
alibaba.wdk.marketing.coupon.queryitems

查询优惠券活动下面的商品
*/
type AlibabaWdkMarketingCouponQueryitemsAPIRequest struct {
    model.Params
    // 查询入参
    _param   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingCouponQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingCouponQueryitemsRequest() *AlibabaWdkMarketingCouponQueryitemsAPIRequest{
    return &AlibabaWdkMarketingCouponQueryitemsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.queryitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询入参
func (r *AlibabaWdkMarketingCouponQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingCouponQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
    return r._param
}
