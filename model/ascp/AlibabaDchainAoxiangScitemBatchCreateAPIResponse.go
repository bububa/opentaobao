package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemBatchCreateAPIResponse 新建货品 API返回值
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
type AlibabaDchainAoxiangScitemBatchCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangScitemBatchCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemBatchCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangScitemBatchCreateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangScitemBatchCreateAPIResponseModel is 新建货品 成功返回结果
type AlibabaDchainAoxiangScitemBatchCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_batch_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchCreateScitemResponse *BatchCreateScItemResponse `json:"batch_create_scitem_response,omitempty" xml:"batch_create_scitem_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemBatchCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchCreateScitemResponse = nil
}

var poolAlibabaDchainAoxiangScitemBatchCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangScitemBatchCreateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangScitemBatchCreateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangScitemBatchCreateAPIResponse
func GetAlibabaDchainAoxiangScitemBatchCreateAPIResponse() *AlibabaDchainAoxiangScitemBatchCreateAPIResponse {
	return poolAlibabaDchainAoxiangScitemBatchCreateAPIResponse.Get().(*AlibabaDchainAoxiangScitemBatchCreateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangScitemBatchCreateAPIResponse 将 AlibabaDchainAoxiangScitemBatchCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemBatchCreateAPIResponse(v *AlibabaDchainAoxiangScitemBatchCreateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemBatchCreateAPIResponse.Put(v)
}
