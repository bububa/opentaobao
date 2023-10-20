package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest ERP全量同步销售库存数量 API请求
// alibaba.dchain.aoxiang.channel.inventory.batch.upload
//
// ERP全量同步销售库存数量
type AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest struct {
	model.Params
	// 上传渠道库存量入参
	_batchUploadChannelInventoryRequest *BatchUploadChannelInventoryRequest
}

// NewAlibabadchainaoxiangchannelinventorybatchuploadRequest 初始化AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest对象
func NewAlibabadchainaoxiangchannelinventorybatchuploadRequest() *AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest {
	return &AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.channel.inventory.batch.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchUploadChannelInventoryRequest is BatchUploadChannelInventoryRequest Setter
// 上传渠道库存量入参
func (r *AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest) SetBatchUploadChannelInventoryRequest(_batchUploadChannelInventoryRequest *BatchUploadChannelInventoryRequest) error {
	r._batchUploadChannelInventoryRequest = _batchUploadChannelInventoryRequest
	r.Set("batch_upload_channel_inventory_request", _batchUploadChannelInventoryRequest)
	return nil
}

// GetBatchUploadChannelInventoryRequest BatchUploadChannelInventoryRequest Getter
func (r AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest) GetBatchUploadChannelInventoryRequest() *BatchUploadChannelInventoryRequest {
	return r._batchUploadChannelInventoryRequest
}
