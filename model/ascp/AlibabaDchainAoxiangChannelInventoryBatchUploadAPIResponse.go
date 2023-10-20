package ascp

import (
	"encoding/xml"

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

// AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel is ERP全量同步销售库存数量 成功返回结果
type AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_channel_inventory_batch_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchUploadChannelInventoryResponse *BatchUploadChannelInventoryResponse `json:"batch_upload_channel_inventory_response,omitempty" xml:"batch_upload_channel_inventory_response,omitempty"`
}
