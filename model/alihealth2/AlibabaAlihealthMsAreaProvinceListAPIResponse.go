package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMsAreaProvinceListAPIResponse 疫苗预约省份列表查询 API返回值
// alibaba.alihealth.ms.area.province.list
//
// 微信小程序疫苗预约省份列表查询
type AlibabaAlihealthMsAreaProvinceListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMsAreaProvinceListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMsAreaProvinceListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMsAreaProvinceListAPIResponseModel).Reset()
}

// AlibabaAlihealthMsAreaProvinceListAPIResponseModel is 疫苗预约省份列表查询 成功返回结果
type AlibabaAlihealthMsAreaProvinceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_ms_area_province_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMsAreaProvinceListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMsAreaProvinceListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMsAreaProvinceListAPIResponse)
	},
}

// GetAlibabaAlihealthMsAreaProvinceListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMsAreaProvinceListAPIResponse
func GetAlibabaAlihealthMsAreaProvinceListAPIResponse() *AlibabaAlihealthMsAreaProvinceListAPIResponse {
	return poolAlibabaAlihealthMsAreaProvinceListAPIResponse.Get().(*AlibabaAlihealthMsAreaProvinceListAPIResponse)
}

// ReleaseAlibabaAlihealthMsAreaProvinceListAPIResponse 将 AlibabaAlihealthMsAreaProvinceListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMsAreaProvinceListAPIResponse(v *AlibabaAlihealthMsAreaProvinceListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMsAreaProvinceListAPIResponse.Put(v)
}
