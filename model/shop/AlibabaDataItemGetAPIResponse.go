package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDataItemGetAPIResponse 获取商品 API返回值
// alibaba.data.item.get
//
// 获取商品信息，作为客户端Weex鉴权的虚拟api
type AlibabaDataItemGetAPIResponse struct {
	model.CommonResponse
	AlibabaDataItemGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDataItemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDataItemGetAPIResponseModel).Reset()
}

// AlibabaDataItemGetAPIResponseModel is 获取商品 成功返回结果
type AlibabaDataItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_data_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取商品信息，作为客户端Weex鉴权的虚拟api
	Unnamed string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDataItemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Unnamed = ""
}

var poolAlibabaDataItemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDataItemGetAPIResponse)
	},
}

// GetAlibabaDataItemGetAPIResponse 从 sync.Pool 获取 AlibabaDataItemGetAPIResponse
func GetAlibabaDataItemGetAPIResponse() *AlibabaDataItemGetAPIResponse {
	return poolAlibabaDataItemGetAPIResponse.Get().(*AlibabaDataItemGetAPIResponse)
}

// ReleaseAlibabaDataItemGetAPIResponse 将 AlibabaDataItemGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDataItemGetAPIResponse(v *AlibabaDataItemGetAPIResponse) {
	v.Reset()
	poolAlibabaDataItemGetAPIResponse.Put(v)
}
