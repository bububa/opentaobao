package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心门店商品修改 API返回值 
alibaba.wdk.item.storesku.update

五道口商品中心门店商品修改
*/
type AlibabaWdkItemStoreskuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemStoreskuUpdateResponse
}

// 五道口商品中心门店商品修改 成功返回结果
type AlibabaWdkItemStoreskuUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_storesku_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkItemStoreskuUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
