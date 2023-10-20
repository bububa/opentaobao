package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse 更新组合货品 API返回值
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
type AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel is 更新组合货品 成功返回结果
type AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_combinescitem_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchUpdateCombineScitemResponse *BatchUpdateCombineScItemResponse `json:"batch_update_combine_scitem_response,omitempty" xml:"batch_update_combine_scitem_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchUpdateCombineScitemResponse = nil
}

var poolAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse
func GetAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse() *AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse {
	return poolAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse.Get().(*AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse 将 AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse(v *AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse.Put(v)
}
