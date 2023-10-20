package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticssellerwriteoffAPIResponse 商家配送核销 API返回值
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
type AlibabaascplogisticssellerwriteoffAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticssellerwriteoffAPIResponseModel
}

// AlibabaascplogisticssellerwriteoffAPIResponseModel is 商家配送核销 成功返回结果
type AlibabaascplogisticssellerwriteoffAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_writeoff_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
