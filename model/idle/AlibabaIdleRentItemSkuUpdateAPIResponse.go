package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemSkuUpdateAPIResponse 更新/增加sku信息 API返回值
// alibaba.idle.rent.item.sku.update
//
// 更新/增加sku信息
type AlibabaIdleRentItemSkuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentItemSkuUpdateAPIResponseModel
}

// AlibabaIdleRentItemSkuUpdateAPIResponseModel is 更新/增加sku信息 成功返回结果
type AlibabaIdleRentItemSkuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_item_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentItemSkuUpdateTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
