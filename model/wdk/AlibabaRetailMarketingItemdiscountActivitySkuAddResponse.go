package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
特价活动新增商品 APIResponse
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_retail_marketing_itemdiscount_activity_sku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"