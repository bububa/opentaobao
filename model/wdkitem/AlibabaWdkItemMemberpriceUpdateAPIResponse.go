package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品售价会员价修改 API返回值 
alibaba.wdk.item.memberprice.update

商品售价会员价修改
*/
type AlibabaWdkItemMemberpriceUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMemberpriceUpdateAPIResponseModel
}

// 商品售价会员价修改 成功返回结果
type AlibabaWdkItemMemberpriceUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_memberprice_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkItemMemberpriceUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
