package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest ERP全量同步销售库存数量 API请求
// alibaba.dchain.aoxiang.channel.inventory.batch.upload
//
// ERP全量同步销售库存数量
type AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest struct {
	model.Params
	// 上传渠道库存量入参
	_batchUploadChannelInventoryRequest *BatchUploadChannelInventoryRequest
}

// NewAlibabaDchainAoxiangChannelInventoryBatchUploadRequest 初始化AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest对象
func NewAlibabaDchainAoxiangChannelInventoryBatchUploadRequest() *AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest {
	return &AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.channel.inventory.batch.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchUploadChannelInventoryRequest is BatchUploadChannelInventoryRequest Setter
// 上传渠道库存量入参
func (r *AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest) SetBatchUploadChannelInventoryRequest(_batchUploadChannelInventoryRequest *BatchUploadChannelInventoryRequest) error {
	r._batchUploadChannelInventoryRequest = _batchUploadChannelInventoryRequest
	r.Set("batch_upload_channel_inventory_request", _batchUploadChannelInventoryRequest)
	return nil
}

// GetBatchUploadChannelInventoryRequest BatchUploadChannelInventoryRequest Getter
func (r AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest) GetBatchUploadChannelInventoryRequest() *BatchUploadChannelInventoryRequest {
	return r._batchUploadChannelInventoryRequest
}
