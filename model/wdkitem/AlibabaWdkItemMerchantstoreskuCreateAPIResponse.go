package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantstoreskuCreateAPIResponse
门店商品信息新建 API返回值
alibaba.wdk.item.merchantstoresku.create

门店商品信息新建 */
type AlibabaWdkItemMerchantstoreskuCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel
}

// AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel is 门店商品信息新建 成功返回结果
type AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantstoresku_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantstoreskuCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
