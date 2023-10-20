package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse 发货单批量查询 API返回值
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
type AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangConsignorderBatchQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangConsignorderBatchQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangConsignorderBatchQueryAPIResponseModel is 发货单批量查询 成功返回结果
type AlibabaDchainAoxiangConsignorderBatchQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_batch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchQueryConsignorderResponse *BatchQueryConsignOrderResponse `json:"batch_query_consignorder_response,omitempty" xml:"batch_query_consignorder_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangConsignorderBatchQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchQueryConsignorderResponse = nil
}

var poolAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse
func GetAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse() *AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse {
	return poolAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse.Get().(*AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse 将 AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse(v *AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangConsignorderBatchQueryAPIResponse.Put(v)
}
