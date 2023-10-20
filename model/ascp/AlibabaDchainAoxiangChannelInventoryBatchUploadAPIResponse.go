package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse ERP全量同步销售库存数量 API返回值
// alibaba.dchain.aoxiang.channel.inventory.batch.upload
//
// ERP全量同步销售库存数量
type AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel is ERP全量同步销售库存数量 成功返回结果
type AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_channel_inventory_batch_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchUploadChannelInventoryResponse *BatchUploadChannelInventoryResponse `json:"batch_upload_channel_inventory_response,omitempty" xml:"batch_upload_channel_inventory_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchUploadChannelInventoryResponse = nil
}

var poolAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse)
	},
}

// GetAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse
func GetAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse() *AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse {
	return poolAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse.Get().(*AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse)
}

// ReleaseAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse 将 AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse(v *AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse.Put(v)
}
