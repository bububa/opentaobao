package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKoubeishopsPropertyGetAPIResponse 口碑店铺列表推荐 API返回值
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
type AlibabaKoubeishopsPropertyGetAPIResponse struct {
	model.CommonResponse
	AlibabaKoubeishopsPropertyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKoubeishopsPropertyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKoubeishopsPropertyGetAPIResponseModel).Reset()
}

// AlibabaKoubeishopsPropertyGetAPIResponseModel is 口碑店铺列表推荐 成功返回结果
type AlibabaKoubeishopsPropertyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_koubeishops_property_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenApiSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKoubeishopsPropertyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKoubeishopsPropertyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKoubeishopsPropertyGetAPIResponse)
	},
}

// GetAlibabaKoubeishopsPropertyGetAPIResponse 从 sync.Pool 获取 AlibabaKoubeishopsPropertyGetAPIResponse
func GetAlibabaKoubeishopsPropertyGetAPIResponse() *AlibabaKoubeishopsPropertyGetAPIResponse {
	return poolAlibabaKoubeishopsPropertyGetAPIResponse.Get().(*AlibabaKoubeishopsPropertyGetAPIResponse)
}

// ReleaseAlibabaKoubeishopsPropertyGetAPIResponse 将 AlibabaKoubeishopsPropertyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKoubeishopsPropertyGetAPIResponse(v *AlibabaKoubeishopsPropertyGetAPIResponse) {
	v.Reset()
	poolAlibabaKoubeishopsPropertyGetAPIResponse.Put(v)
}
