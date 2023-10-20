package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticssellersendAPIResponse 商家配送发货 API返回值
// alibaba.ascp.logistics.seller.send
//
// 该API提供商家配送发货能力，使用该接口发货，交易订单状态会直接变成卖家已发货
type AlibabaascplogisticssellersendAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticssellersendAPIResponseModel
}

// AlibabaascplogisticssellersendAPIResponseModel is 商家配送发货 成功返回结果
type AlibabaascplogisticssellersendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
