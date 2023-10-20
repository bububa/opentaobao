package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse 商家关系查询分销商 API返回值
// alibaba.dchain.aoxiang.cooperate.distributor.query
//
// 商家关系查询分销商
type AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangCooperateDistributorQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangCooperateDistributorQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangCooperateDistributorQueryAPIResponseModel is 商家关系查询分销商 成功返回结果
type AlibabaDchainAoxiangCooperateDistributorQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_cooperate_distributor_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *TopResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCooperateDistributorQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse
func GetAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse() *AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse {
	return poolAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse.Get().(*AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse 将 AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse(v *AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangCooperateDistributorQueryAPIResponse.Put(v)
}
