package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpSalecategoryQueryAPIResponse 货品品类查询 API返回值
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
type AlibabaAscpSalecategoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpSalecategoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpSalecategoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpSalecategoryQueryAPIResponseModel).Reset()
}

// AlibabaAscpSalecategoryQueryAPIResponseModel is 货品品类查询 成功返回结果
type AlibabaAscpSalecategoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_salecategory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 货品查询结构化对象
	DataList []SalecategoryQueryResponse `json:"data_list,omitempty" xml:"data_list>salecategory_query_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpSalecategoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
}

var poolAlibabaAscpSalecategoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpSalecategoryQueryAPIResponse)
	},
}

// GetAlibabaAscpSalecategoryQueryAPIResponse 从 sync.Pool 获取 AlibabaAscpSalecategoryQueryAPIResponse
func GetAlibabaAscpSalecategoryQueryAPIResponse() *AlibabaAscpSalecategoryQueryAPIResponse {
	return poolAlibabaAscpSalecategoryQueryAPIResponse.Get().(*AlibabaAscpSalecategoryQueryAPIResponse)
}

// ReleaseAlibabaAscpSalecategoryQueryAPIResponse 将 AlibabaAscpSalecategoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpSalecategoryQueryAPIResponse(v *AlibabaAscpSalecategoryQueryAPIResponse) {
	v.Reset()
	poolAlibabaAscpSalecategoryQueryAPIResponse.Put(v)
}
