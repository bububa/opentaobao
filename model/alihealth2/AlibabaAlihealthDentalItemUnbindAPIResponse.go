package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemUnbindAPIResponse ISV解绑商品 API返回值
// alibaba.alihealth.dental.item.unbind
//
// ISV解绑商品
type AlibabaAlihealthDentalItemUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalItemUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalItemUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalItemUnbindAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalItemUnbindAPIResponseModel is ISV解绑商品 成功返回结果
type AlibabaAlihealthDentalItemUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_item_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalItemUnbindMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalItemUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalItemUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalItemUnbindAPIResponse)
	},
}

// GetAlibabaAlihealthDentalItemUnbindAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalItemUnbindAPIResponse
func GetAlibabaAlihealthDentalItemUnbindAPIResponse() *AlibabaAlihealthDentalItemUnbindAPIResponse {
	return poolAlibabaAlihealthDentalItemUnbindAPIResponse.Get().(*AlibabaAlihealthDentalItemUnbindAPIResponse)
}

// ReleaseAlibabaAlihealthDentalItemUnbindAPIResponse 将 AlibabaAlihealthDentalItemUnbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalItemUnbindAPIResponse(v *AlibabaAlihealthDentalItemUnbindAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalItemUnbindAPIResponse.Put(v)
}
