package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品 APIResponse
alibaba.wdk.merchant.storeitem.update

修改门店商品
*/
type AlibabaWdkMerchantStoreitemUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_merchant_storeitem_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkMerchantStoreitemUpdateResult `json:"result,omitempty" xml:"