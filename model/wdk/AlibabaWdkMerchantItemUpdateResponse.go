package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商家商品 APIResponse
alibaba.wdk.merchant.item.update

修改商家商品
*/
type AlibabaWdkMerchantItemUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_merchant_item_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkMerchantItemUpdateResult `json:"result,omitempty" xml:"