package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品删除 APIResponse
alibaba.wdk.coupon.sku.remove

优惠券商品删除
*/
type AlibabaWdkCouponSkuRemoveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_coupon_sku_remove_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponSkuRemoveApiResult `json:"result,omitempty" xml:"