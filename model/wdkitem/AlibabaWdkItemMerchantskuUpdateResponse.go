package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品修改 APIResponse
alibaba.wdk.item.merchantsku.update

商家商品修改
*/
type AlibabaWdkItemMerchantskuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMerchantskuUpdateResponse
}

type AlibabaWdkItemMerchantskuUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemMerchantskuUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
