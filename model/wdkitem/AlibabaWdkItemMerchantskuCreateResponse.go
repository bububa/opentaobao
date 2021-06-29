package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品信息新建 APIResponse
alibaba.wdk.item.merchantsku.create

商家商品信息新建
*/
type AlibabaWdkItemMerchantskuCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMerchantskuCreateResponse
}

type AlibabaWdkItemMerchantskuCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemMerchantskuCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
