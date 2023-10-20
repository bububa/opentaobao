package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemBindAPIResponse ISV绑定外部门店id和外部商品id API返回值
// alibaba.alihealth.dental.item.bind
//
// ISV绑定外部门店id和外部商品id
type AlibabaAlihealthDentalItemBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalItemBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalItemBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalItemBindAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalItemBindAPIResponseModel is ISV绑定外部门店id和外部商品id 成功返回结果
type AlibabaAlihealthDentalItemBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_item_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalItemBindMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalItemBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalItemBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalItemBindAPIResponse)
	},
}

// GetAlibabaAlihealthDentalItemBindAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalItemBindAPIResponse
func GetAlibabaAlihealthDentalItemBindAPIResponse() *AlibabaAlihealthDentalItemBindAPIResponse {
	return poolAlibabaAlihealthDentalItemBindAPIResponse.Get().(*AlibabaAlihealthDentalItemBindAPIResponse)
}

// ReleaseAlibabaAlihealthDentalItemBindAPIResponse 将 AlibabaAlihealthDentalItemBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalItemBindAPIResponse(v *AlibabaAlihealthDentalItemBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalItemBindAPIResponse.Put(v)
}
