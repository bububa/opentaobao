package newretail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItApAddressSetAPIResponse setApAddressNew API返回值
// alibaba.it.ap.address.set
//
// 该接口可为ISV系统提供 ap位置信息维护的功能
type AlibabaItApAddressSetAPIResponse struct {
	model.CommonResponse
	AlibabaItApAddressSetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItApAddressSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItApAddressSetAPIResponseModel).Reset()
}

// AlibabaItApAddressSetAPIResponseModel is setApAddressNew 成功返回结果
type AlibabaItApAddressSetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_ap_address_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaItApAddressSetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItApAddressSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaItApAddressSetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItApAddressSetAPIResponse)
	},
}

// GetAlibabaItApAddressSetAPIResponse 从 sync.Pool 获取 AlibabaItApAddressSetAPIResponse
func GetAlibabaItApAddressSetAPIResponse() *AlibabaItApAddressSetAPIResponse {
	return poolAlibabaItApAddressSetAPIResponse.Get().(*AlibabaItApAddressSetAPIResponse)
}

// ReleaseAlibabaItApAddressSetAPIResponse 将 AlibabaItApAddressSetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItApAddressSetAPIResponse(v *AlibabaItApAddressSetAPIResponse) {
	v.Reset()
	poolAlibabaItApAddressSetAPIResponse.Put(v)
}
