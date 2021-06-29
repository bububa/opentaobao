package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品售价会员价修改 APIResponse
alibaba.wdk.item.memberprice.update

商品售价会员价修改
*/
type AlibabaWdkItemMemberpriceUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMemberpriceUpdateResponse
}

type AlibabaWdkItemMemberpriceUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_memberprice_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemMemberpriceUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
