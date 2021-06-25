package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动下的商品 APIRequest
alibaba.wdk.marketing.coupon.queryitems

查询优惠券活动下面的商品
*/
type AlibabaWdkMarketingCouponQueryitemsRequest struct {
    model.Params

    // 查询入参
    param   *ActivitySkuQuery 

}

func NewAlibabaWdkMarketingCouponQueryitemsRequest() *AlibabaWdkMarketingCouponQueryitemsRequest{
    return &AlibabaWdkMarketingCouponQueryitemsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingCouponQueryitemsRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.queryitems"
}

func (r AlibabaWdkMarketingCouponQueryitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingCouponQueryitemsRequest) SetParam(param *ActivitySkuQuery) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingCouponQueryitemsRequest) GetParam() *ActivitySkuQuery {
    return r.param
}

