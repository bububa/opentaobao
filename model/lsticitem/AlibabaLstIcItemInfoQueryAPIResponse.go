package lsticitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstIcItemInfoQueryAPIResponse 商品信息查询 API返回值
// alibaba.lst.ic.item.info.query
//
// 查询商品信息
type AlibabaLstIcItemInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstIcItemInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstIcItemInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstIcItemInfoQueryAPIResponseModel).Reset()
}

// AlibabaLstIcItemInfoQueryAPIResponseModel is 商品信息查询 成功返回结果
type AlibabaLstIcItemInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_ic_item_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstIcItemInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstIcItemInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstIcItemInfoQueryAPIResponse)
	},
}

// GetAlibabaLstIcItemInfoQueryAPIResponse 从 sync.Pool 获取 AlibabaLstIcItemInfoQueryAPIResponse
func GetAlibabaLstIcItemInfoQueryAPIResponse() *AlibabaLstIcItemInfoQueryAPIResponse {
	return poolAlibabaLstIcItemInfoQueryAPIResponse.Get().(*AlibabaLstIcItemInfoQueryAPIResponse)
}

// ReleaseAlibabaLstIcItemInfoQueryAPIResponse 将 AlibabaLstIcItemInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstIcItemInfoQueryAPIResponse(v *AlibabaLstIcItemInfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstIcItemInfoQueryAPIResponse.Put(v)
}
