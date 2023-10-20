package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemQueryAPIResponse 货品查询 API返回值
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
type AlibabaDchainAoxiangScitemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangScitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangScitemQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangScitemQueryAPIResponseModel is 货品查询 成功返回结果
type AlibabaDchainAoxiangScitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	QueryScitemResponse *TopResponse `json:"query_scitem_response,omitempty" xml:"query_scitem_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QueryScitemResponse = nil
}

var poolAlibabaDchainAoxiangScitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangScitemQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangScitemQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangScitemQueryAPIResponse
func GetAlibabaDchainAoxiangScitemQueryAPIResponse() *AlibabaDchainAoxiangScitemQueryAPIResponse {
	return poolAlibabaDchainAoxiangScitemQueryAPIResponse.Get().(*AlibabaDchainAoxiangScitemQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangScitemQueryAPIResponse 将 AlibabaDchainAoxiangScitemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemQueryAPIResponse(v *AlibabaDchainAoxiangScitemQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemQueryAPIResponse.Put(v)
}
