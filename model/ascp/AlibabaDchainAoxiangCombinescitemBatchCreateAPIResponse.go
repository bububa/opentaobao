package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse 新建组合货品 API返回值
// alibaba.dchain.aoxiang.combinescitem.batch.create
//
// 新建组合货品
type AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponseModel is 新建组合货品 成功返回结果
type AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_combinescitem_batch_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchCreateCombineScitemResponse *BatchCreateCombineScItemResponse `json:"batch_create_combine_scitem_response,omitempty" xml:"batch_create_combine_scitem_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchCreateCombineScitemResponse = nil
}

var poolAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse
func GetAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse() *AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse {
	return poolAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse.Get().(*AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse 将 AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse(v *AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse.Put(v)
}
