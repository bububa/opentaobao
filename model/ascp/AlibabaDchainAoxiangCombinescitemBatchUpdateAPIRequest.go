package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest 更新组合货品 API请求
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
type AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest struct {
	model.Params
	// 批量更新组合货品入参
	_batchUpdateCombineScitemRequest *BatchUpdateCombineScItemRequest
}

// NewAlibabaDchainAoxiangCombinescitemBatchUpdateRequest 初始化AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest对象
func NewAlibabaDchainAoxiangCombinescitemBatchUpdateRequest() *AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest {
	return &AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.combinescitem.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchUpdateCombineScitemRequest is BatchUpdateCombineScitemRequest Setter
// 批量更新组合货品入参
func (r *AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest) SetBatchUpdateCombineScitemRequest(_batchUpdateCombineScitemRequest *BatchUpdateCombineScItemRequest) error {
	r._batchUpdateCombineScitemRequest = _batchUpdateCombineScitemRequest
	r.Set("batch_update_combine_scitem_request", _batchUpdateCombineScitemRequest)
	return nil
}

// GetBatchUpdateCombineScitemRequest BatchUpdateCombineScitemRequest Getter
func (r AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest) GetBatchUpdateCombineScitemRequest() *BatchUpdateCombineScItemRequest {
	return r._batchUpdateCombineScitemRequest
}
