package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductGroupGetAPIResponse 查询指定产品分组的下一层子分组 API返回值
// alibaba.scbp.product.group.get
//
// 查询指定产品分组的下一层子分组
type AlibabaScbpProductGroupGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpProductGroupGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpProductGroupGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpProductGroupGetAPIResponseModel).Reset()
}

// AlibabaScbpProductGroupGetAPIResponseModel is 查询指定产品分组的下一层子分组 成功返回结果
type AlibabaScbpProductGroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_product_group_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下一层分组列表
	RoductGroupList []TopProductGroupDto `json:"roduct_group_list,omitempty" xml:"roduct_group_list>top_product_group_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpProductGroupGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RoductGroupList = m.RoductGroupList[:0]
}

var poolAlibabaScbpProductGroupGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpProductGroupGetAPIResponse)
	},
}

// GetAlibabaScbpProductGroupGetAPIResponse 从 sync.Pool 获取 AlibabaScbpProductGroupGetAPIResponse
func GetAlibabaScbpProductGroupGetAPIResponse() *AlibabaScbpProductGroupGetAPIResponse {
	return poolAlibabaScbpProductGroupGetAPIResponse.Get().(*AlibabaScbpProductGroupGetAPIResponse)
}

// ReleaseAlibabaScbpProductGroupGetAPIResponse 将 AlibabaScbpProductGroupGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpProductGroupGetAPIResponse(v *AlibabaScbpProductGroupGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpProductGroupGetAPIResponse.Put(v)
}
