package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalBindshopanditemAPIResponse 绑定门店信息，商品信息 API返回值
// alibaba.alihealth.reserve.dental.bindshopanditem
//
// 绑定门店信息，商品信息
type AlibabaAlihealthReserveDentalBindshopanditemAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReserveDentalBindshopanditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthReserveDentalBindshopanditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthReserveDentalBindshopanditemAPIResponseModel).Reset()
}

// AlibabaAlihealthReserveDentalBindshopanditemAPIResponseModel is 绑定门店信息，商品信息 成功返回结果
type AlibabaAlihealthReserveDentalBindshopanditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_bindshopanditem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthReserveDentalBindshopanditemResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthReserveDentalBindshopanditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthReserveDentalBindshopanditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthReserveDentalBindshopanditemAPIResponse)
	},
}

// GetAlibabaAlihealthReserveDentalBindshopanditemAPIResponse 从 sync.Pool 获取 AlibabaAlihealthReserveDentalBindshopanditemAPIResponse
func GetAlibabaAlihealthReserveDentalBindshopanditemAPIResponse() *AlibabaAlihealthReserveDentalBindshopanditemAPIResponse {
	return poolAlibabaAlihealthReserveDentalBindshopanditemAPIResponse.Get().(*AlibabaAlihealthReserveDentalBindshopanditemAPIResponse)
}

// ReleaseAlibabaAlihealthReserveDentalBindshopanditemAPIResponse 将 AlibabaAlihealthReserveDentalBindshopanditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthReserveDentalBindshopanditemAPIResponse(v *AlibabaAlihealthReserveDentalBindshopanditemAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthReserveDentalBindshopanditemAPIResponse.Put(v)
}
