package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantskuupdateAPIResponse 商家商品修改 API返回值
// alibaba.wdk.item.merchantsku.update
//
// 商家商品修改
type AlibabawdkitemmerchantskuupdateAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemmerchantskuupdateAPIResponseModel
}

// AlibabawdkitemmerchantskuupdateAPIResponseModel is 商家商品修改 成功返回结果
type AlibabawdkitemmerchantskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkitemmerchantskuupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
