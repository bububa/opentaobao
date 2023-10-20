package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemListAPIResponse ISV获取口腔标品列表 API返回值
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
type AlibabaAlihealthDentalItemListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalItemListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalItemListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalItemListAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalItemListAPIResponseModel is ISV获取口腔标品列表 成功返回结果
type AlibabaAlihealthDentalItemListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_item_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalItemListMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalItemListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalItemListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalItemListAPIResponse)
	},
}

// GetAlibabaAlihealthDentalItemListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalItemListAPIResponse
func GetAlibabaAlihealthDentalItemListAPIResponse() *AlibabaAlihealthDentalItemListAPIResponse {
	return poolAlibabaAlihealthDentalItemListAPIResponse.Get().(*AlibabaAlihealthDentalItemListAPIResponse)
}

// ReleaseAlibabaAlihealthDentalItemListAPIResponse 将 AlibabaAlihealthDentalItemListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalItemListAPIResponse(v *AlibabaAlihealthDentalItemListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalItemListAPIResponse.Put(v)
}
