package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.delete

同城零售商品池活动删除
*/
type AlibabaRetailMarketingItempoolActivityDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_retail_marketing_itempool_activity_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"