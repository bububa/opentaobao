package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantItemUpdateAPIResponse 修改商家商品 API返回值
// alibaba.wdk.merchant.item.update
//
// 修改商家商品
type AlibabaWdkMerchantItemUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantItemUpdateAPIResponseModel
}

// AlibabaWdkMerchantItemUpdateAPIResponseModel is 修改商家商品 成功返回结果
type AlibabaWdkMerchantItemUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkMerchantItemUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
