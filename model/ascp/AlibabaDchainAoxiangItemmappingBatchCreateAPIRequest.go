package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest 新建商货品关联 API请求
// alibaba.dchain.aoxiang.itemmapping.batch.create
//
// 新建商货品关联
type AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest struct {
	model.Params
	// 创建/更新商货品关联入参
	_batchCreateItemMappingRequest *BatchCreateItemMappingRequest
}

// NewAlibabaDchainAoxiangItemmappingBatchCreateRequest 初始化AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest对象
func NewAlibabaDchainAoxiangItemmappingBatchCreateRequest() *AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest {
	return &AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.batch.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchCreateItemMappingRequest is BatchCreateItemMappingRequest Setter
// 创建/更新商货品关联入参
func (r *AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest) SetBatchCreateItemMappingRequest(_batchCreateItemMappingRequest *BatchCreateItemMappingRequest) error {
	r._batchCreateItemMappingRequest = _batchCreateItemMappingRequest
	r.Set("batch_create_item_mapping_request", _batchCreateItemMappingRequest)
	return nil
}

// GetBatchCreateItemMappingRequest BatchCreateItemMappingRequest Getter
func (r AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest) GetBatchCreateItemMappingRequest() *BatchCreateItemMappingRequest {
	return r._batchCreateItemMappingRequest
}
