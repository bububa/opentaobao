package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品一品多码维护操作 APIResponse
alibaba.wdk.item.morebarcode.ops

商品一品多码维护操作
*/
type AlibabaWdkItemMorebarcodeOpsAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMorebarcodeOpsResponse
}

type AlibabaWdkItemMorebarcodeOpsResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_morebarcode_ops_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemMorebarcodeOpsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
