package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过券模板查询券活动id接口 API请求
alibaba.wdk.coupon.template.queryumpactid

当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口
*/
type AlibabaWdkCouponTemplateQueryumpactidAPIRequest struct {
    model.Params
    // 券模版id列表
    _sourceIds   []int64
    // 优惠券类型
    _wdkCouponType   int64
}

// 初始化AlibabaWdkCouponTemplateQueryumpactidAPIRequest对象
func NewAlibabaWdkCouponTemplateQueryumpactidRequest() *AlibabaWdkCouponTemplateQueryumpactidAPIRequest{
    return &AlibabaWdkCouponTemplateQueryumpactidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.template.queryumpactid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SourceIds Setter
// 券模版id列表
func (r *AlibabaWdkCouponTemplateQueryumpactidAPIRequest) SetSourceIds(_sourceIds []int64) error {
    r._sourceIds = _sourceIds
    r.Set("source_ids", _sourceIds)
    return nil
}

// SourceIds Getter
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetSourceIds() []int64 {
    return r._sourceIds
}
// WdkCouponType Setter
// 优惠券类型
func (r *AlibabaWdkCouponTemplateQueryumpactidAPIRequest) SetWdkCouponType(_wdkCouponType int64) error {
    r._wdkCouponType = _wdkCouponType
    r.Set("wdk_coupon_type", _wdkCouponType)
    return nil
}

// WdkCouponType Getter
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetWdkCouponType() int64 {
    return r._wdkCouponType
}
