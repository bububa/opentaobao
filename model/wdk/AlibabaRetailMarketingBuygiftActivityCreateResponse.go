package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.create

同城供给买赠活动创建
*/
type AlibabaRetailMarketingBuygiftActivityCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_retail_marketing_buygift_activity_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"