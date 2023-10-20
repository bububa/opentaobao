package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalUnbinditemAPIResponse 解绑商品信息 API返回值
// alibaba.alihealth.reserve.dental.unbinditem
//
// 绑定门店信息，商品信息
type AlibabaAlihealthReserveDentalUnbinditemAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReserveDentalUnbinditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthReserveDentalUnbinditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthReserveDentalUnbinditemAPIResponseModel).Reset()
}

// AlibabaAlihealthReserveDentalUnbinditemAPIResponseModel is 解绑商品信息 成功返回结果
type AlibabaAlihealthReserveDentalUnbinditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_unbinditem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthReserveDentalUnbinditemResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthReserveDentalUnbinditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthReserveDentalUnbinditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthReserveDentalUnbinditemAPIResponse)
	},
}

// GetAlibabaAlihealthReserveDentalUnbinditemAPIResponse 从 sync.Pool 获取 AlibabaAlihealthReserveDentalUnbinditemAPIResponse
func GetAlibabaAlihealthReserveDentalUnbinditemAPIResponse() *AlibabaAlihealthReserveDentalUnbinditemAPIResponse {
	return poolAlibabaAlihealthReserveDentalUnbinditemAPIResponse.Get().(*AlibabaAlihealthReserveDentalUnbinditemAPIResponse)
}

// ReleaseAlibabaAlihealthReserveDentalUnbinditemAPIResponse 将 AlibabaAlihealthReserveDentalUnbinditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthReserveDentalUnbinditemAPIResponse(v *AlibabaAlihealthReserveDentalUnbinditemAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthReserveDentalUnbinditemAPIResponse.Put(v)
}
