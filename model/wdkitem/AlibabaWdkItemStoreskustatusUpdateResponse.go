package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品状态 APIResponse
alibaba.wdk.item.storeskustatus.update

五道口商品 修改门店商品状态
*/
type AlibabaWdkItemStoreskustatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemStoreskustatusUpdateResponse
}

type AlibabaWdkItemStoreskustatusUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_storeskustatus_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemStoreskustatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
