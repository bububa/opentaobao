package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过券模板查询券活动id接口 APIRequest
alibaba.wdk.coupon.template.queryumpactid

当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口
*/
type AlibabaWdkCouponTemplateQueryumpactidRequest struct {
    model.Params

    // 券模版id列表
    sourceIds   []int64 

    // 优惠券类型
    wdkCouponType   int64 

}

func NewAlibabaWdkCouponTemplateQueryumpactidRequest() *AlibabaWdkCouponTemplateQueryumpactidRequest{
    return &AlibabaWdkCouponTemplateQueryumpactidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponTemplateQueryumpactidRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.queryumpactid"
}

func (r AlibabaWdkCouponTemplateQueryumpactidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponTemplateQueryumpactidRequest) SetSourceIds(sourceIds []int64) error {
    r.sourceIds = sourceIds
    r.Set("source_ids", sourceIds)
    return nil
}

func (r AlibabaWdkCouponTemplateQueryumpactidRequest) GetSourceIds() []int64 {
    return r.sourceIds
}

func (r *AlibabaWdkCouponTemplateQueryumpactidRequest) SetWdkCouponType(wdkCouponType int64) error {
    r.wdkCouponType = wdkCouponType
    r.Set("wdk_coupon_type", wdkCouponType)
    return nil
}

func (r AlibabaWdkCouponTemplateQueryumpactidRequest) GetWdkCouponType() int64 {
    return r.wdkCouponType
}

