package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品一品多码维护操作 API返回值 
alibaba.wdk.item.morebarcode.ops

商品一品多码维护操作
*/
type AlibabaWdkItemMorebarcodeOpsAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMorebarcodeOpsResponse
}

// 商品一品多码维护操作 成功返回结果
type AlibabaWdkItemMorebarcodeOpsResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_morebarcode_ops_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkItemMorebarcodeOpsResult `json:"result,omitempty" xml:"result,omitempty"`
}
