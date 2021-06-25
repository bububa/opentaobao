package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过券模板查询券活动id接口 APIResponse
alibaba.wdk.coupon.template.queryumpactid

当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口
*/
type AlibabaWdkCouponTemplateQueryumpactidAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponTemplateQueryumpactidResponse `json:"alibaba_wdk_coupon_template_queryumpactid_response,omitempty"`
}

type AlibabaWdkCouponTemplateQueryumpactidResponse struct {

    // 根据站点名称查询产品
    ApiResult   *AlibabaWdkCouponTemplateQueryumpactidApiResult `json:"api_result,omitempty"`

}
