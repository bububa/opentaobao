package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantskucreateAPIResponse 商家商品信息新建 API返回值
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
type AlibabawdkitemmerchantskucreateAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemmerchantskucreateAPIResponseModel
}

// AlibabawdkitemmerchantskucreateAPIResponseModel is 商家商品信息新建 成功返回结果
type AlibabawdkitemmerchantskucreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkitemmerchantskucreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
