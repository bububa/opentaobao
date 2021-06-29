package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心价格修改 API返回值 
alibaba.wdk.item.price.update

修改门店商品售价和会员价
*/
type AlibabaWdkItemPriceUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemPriceUpdateResponse
}

// 五道口商品中心价格修改 成功返回结果
type AlibabaWdkItemPriceUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_price_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // hsfResult
    HsfResult   *AlibabaWdkItemPriceUpdateResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
