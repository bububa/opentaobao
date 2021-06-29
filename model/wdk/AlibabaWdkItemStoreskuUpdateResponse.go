package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心门店商品修改 APIResponse
alibaba.wdk.item.storesku.update

五道口商品中心门店商品修改
*/
type AlibabaWdkItemStoreskuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemStoreskuUpdateResponse
}

type AlibabaWdkItemStoreskuUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_storesku_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemStoreskuUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
