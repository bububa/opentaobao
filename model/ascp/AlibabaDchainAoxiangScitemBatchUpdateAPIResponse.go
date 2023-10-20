package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemBatchUpdateAPIResponse alibaba.dchain.aoxiang.scitem.batch.update API返回值
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
type AlibabaDchainAoxiangScitemBatchUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangScitemBatchUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemBatchUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangScitemBatchUpdateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangScitemBatchUpdateAPIResponseModel is alibaba.dchain.aoxiang.scitem.batch.update 成功返回结果
type AlibabaDchainAoxiangScitemBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchUpdateScitemResponse *BatchUpdateScItemResponse `json:"batch_update_scitem_response,omitempty" xml:"batch_update_scitem_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemBatchUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchUpdateScitemResponse = nil
}

var poolAlibabaDchainAoxiangScitemBatchUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangScitemBatchUpdateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangScitemBatchUpdateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangScitemBatchUpdateAPIResponse
func GetAlibabaDchainAoxiangScitemBatchUpdateAPIResponse() *AlibabaDchainAoxiangScitemBatchUpdateAPIResponse {
	return poolAlibabaDchainAoxiangScitemBatchUpdateAPIResponse.Get().(*AlibabaDchainAoxiangScitemBatchUpdateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangScitemBatchUpdateAPIResponse 将 AlibabaDchainAoxiangScitemBatchUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemBatchUpdateAPIResponse(v *AlibabaDchainAoxiangScitemBatchUpdateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemBatchUpdateAPIResponse.Put(v)
}
