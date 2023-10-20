package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGroupGetAPIResponse 分组信息获取 API返回值
// alibaba.icbu.product.group.get
//
// 分组信息获取
type AlibabaIcbuProductGroupGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductGroupGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductGroupGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductGroupGetAPIResponseModel).Reset()
}

// AlibabaIcbuProductGroupGetAPIResponseModel is 分组信息获取 成功返回结果
type AlibabaIcbuProductGroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_group_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组信息
	ProductGroup *ProductGroup `json:"product_group,omitempty" xml:"product_group,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductGroupGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductGroup = nil
}

var poolAlibabaIcbuProductGroupGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductGroupGetAPIResponse)
	},
}

// GetAlibabaIcbuProductGroupGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductGroupGetAPIResponse
func GetAlibabaIcbuProductGroupGetAPIResponse() *AlibabaIcbuProductGroupGetAPIResponse {
	return poolAlibabaIcbuProductGroupGetAPIResponse.Get().(*AlibabaIcbuProductGroupGetAPIResponse)
}

// ReleaseAlibabaIcbuProductGroupGetAPIResponse 将 AlibabaIcbuProductGroupGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductGroupGetAPIResponse(v *AlibabaIcbuProductGroupGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductGroupGetAPIResponse.Put(v)
}
