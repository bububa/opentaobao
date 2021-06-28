package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rt回传采购价 APIResponse
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价
*/
type AlibabaWdkPurchasePriceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_purchase_price_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // SYSTEM ERROR
    
    ReturnMsg   string `json:"return_msg,omitempty" xml:"