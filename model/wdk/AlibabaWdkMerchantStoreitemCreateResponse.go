package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建门店商品 APIResponse
alibaba.wdk.merchant.storeitem.create

新建门店商品
*/
type AlibabaWdkMerchantStoreitemCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_merchant_storeitem_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    Suc   bool `json:"suc,omitempty" xml:"