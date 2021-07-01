package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantskuCreateAPIResponse
商家商品信息新建 API返回值
alibaba.wdk.item.merchantsku.create

商家商品信息新建 */
type AlibabaWdkItemMerchantskuCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantskuCreateAPIResponseModel
}

// AlibabaWdkItemMerchantskuCreateAPIResponseModel is 商家商品信息新建 成功返回结果
type AlibabaWdkItemMerchantskuCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantskuCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
