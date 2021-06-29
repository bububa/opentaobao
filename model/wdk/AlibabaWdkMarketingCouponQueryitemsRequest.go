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
type AlibabaWdkMarketingCouponQueryitemsRequest struct {
    model.Params
    // 查询入参
    _param   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingCouponQueryitemsRequest对象
func NewAlibabaWdkMarketingCouponQueryitemsRequest() *AlibabaWdkMarketingCouponQueryitemsRequest{
    return &AlibabaWdkMarketingCouponQueryitemsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponQueryitemsRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.queryitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponQueryitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询入参
func (r *AlibabaWdkMarketingCouponQueryitemsRequest) SetParam(_param *ActivitySkuQuery) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingCouponQueryitemsRequest) GetParam() *ActivitySkuQuery {
    return r._param
}
