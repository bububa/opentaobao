package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantstoreskuUpdateAPIResponse 门店商品信息修改 API返回值
// alibaba.wdk.item.merchantstoresku.update
//
// 门店商品信息修改
type AlibabaWdkItemMerchantstoreskuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantstoreskuUpdateAPIResponseModel
}

// AlibabaWdkItemMerchantstoreskuUpdateAPIResponseModel is 门店商品信息修改 成功返回结果
type AlibabaWdkItemMerchantstoreskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantstoresku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantstoreskuUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
