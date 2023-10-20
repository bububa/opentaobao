package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalStoresanditemsAPIResponse 查询商户门店，商品列表 API返回值
// alibaba.alihealth.reserve.dental.storesanditems
//
// 查询商户门店，商品列表
type AlibabaAlihealthReserveDentalStoresanditemsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthReserveDentalStoresanditemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel).Reset()
}

// AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel is 查询商户门店，商品列表 成功返回结果
type AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_storesanditems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthReserveDentalStoresanditemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthReserveDentalStoresanditemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthReserveDentalStoresanditemsAPIResponse)
	},
}

// GetAlibabaAlihealthReserveDentalStoresanditemsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthReserveDentalStoresanditemsAPIResponse
func GetAlibabaAlihealthReserveDentalStoresanditemsAPIResponse() *AlibabaAlihealthReserveDentalStoresanditemsAPIResponse {
	return poolAlibabaAlihealthReserveDentalStoresanditemsAPIResponse.Get().(*AlibabaAlihealthReserveDentalStoresanditemsAPIResponse)
}

// ReleaseAlibabaAlihealthReserveDentalStoresanditemsAPIResponse 将 AlibabaAlihealthReserveDentalStoresanditemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthReserveDentalStoresanditemsAPIResponse(v *AlibabaAlihealthReserveDentalStoresanditemsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthReserveDentalStoresanditemsAPIResponse.Put(v)
}
