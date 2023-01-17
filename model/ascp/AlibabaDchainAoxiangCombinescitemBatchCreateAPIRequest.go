package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest 新建组合货品 API请求
// alibaba.dchain.aoxiang.combinescitem.batch.create
//
// 新建组合货品
type AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest struct {
	model.Params
	// 批量新建组合货品入参
	_batchCreateCombineScitemRequest *BatchCreateCombineScItemRequest
}

// NewAlibabaDchainAoxiangCombinescitemBatchCreateRequest 初始化AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest对象
func NewAlibabaDchainAoxiangCombinescitemBatchCreateRequest() *AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest {
	return &AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.combinescitem.batch.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchCreateCombineScitemRequest is BatchCreateCombineScitemRequest Setter
// 批量新建组合货品入参
func (r *AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest) SetBatchCreateCombineScitemRequest(_batchCreateCombineScitemRequest *BatchCreateCombineScItemRequest) error {
	r._batchCreateCombineScitemRequest = _batchCreateCombineScitemRequest
	r.Set("batch_create_combine_scitem_request", _batchCreateCombineScitemRequest)
	return nil
}

// GetBatchCreateCombineScitemRequest BatchCreateCombineScitemRequest Getter
func (r AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest) GetBatchCreateCombineScitemRequest() *BatchCreateCombineScItemRequest {
	return r._batchCreateCombineScitemRequest
}
