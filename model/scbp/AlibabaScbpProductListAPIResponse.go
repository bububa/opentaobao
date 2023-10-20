package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductListAPIResponse 查询P4P产品 API返回值
// alibaba.scbp.product.list
//
// 查询P4P产品
type AlibabaScbpProductListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpProductListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpProductListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpProductListAPIResponseModel).Reset()
}

// AlibabaScbpProductListAPIResponseModel is 查询P4P产品 成功返回结果
type AlibabaScbpProductListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品列表
	ProductList []TopProductDto `json:"product_list,omitempty" xml:"product_list>top_product_dto,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpProductListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductList = m.ProductList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpProductListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpProductListAPIResponse)
	},
}

// GetAlibabaScbpProductListAPIResponse 从 sync.Pool 获取 AlibabaScbpProductListAPIResponse
func GetAlibabaScbpProductListAPIResponse() *AlibabaScbpProductListAPIResponse {
	return poolAlibabaScbpProductListAPIResponse.Get().(*AlibabaScbpProductListAPIResponse)
}

// ReleaseAlibabaScbpProductListAPIResponse 将 AlibabaScbpProductListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpProductListAPIResponse(v *AlibabaScbpProductListAPIResponse) {
	v.Reset()
	poolAlibabaScbpProductListAPIResponse.Put(v)
}
