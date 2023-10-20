package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeApartmentOuteridAPIResponse 公寓更新outerid API返回值
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
type AlibabaAlihouseNewhomeApartmentOuteridAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeApartmentOuteridAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeApartmentOuteridAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeApartmentOuteridAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeApartmentOuteridAPIResponseModel is 公寓更新outerid 成功返回结果
type AlibabaAlihouseNewhomeApartmentOuteridAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_apartment_outerid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeApartmentOuteridResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeApartmentOuteridAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeApartmentOuteridAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeApartmentOuteridAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeApartmentOuteridAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeApartmentOuteridAPIResponse
func GetAlibabaAlihouseNewhomeApartmentOuteridAPIResponse() *AlibabaAlihouseNewhomeApartmentOuteridAPIResponse {
	return poolAlibabaAlihouseNewhomeApartmentOuteridAPIResponse.Get().(*AlibabaAlihouseNewhomeApartmentOuteridAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeApartmentOuteridAPIResponse 将 AlibabaAlihouseNewhomeApartmentOuteridAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeApartmentOuteridAPIResponse(v *AlibabaAlihouseNewhomeApartmentOuteridAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeApartmentOuteridAPIResponse.Put(v)
}
