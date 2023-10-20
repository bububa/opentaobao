package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangcombinescitembatchupdateAPIRequest 更新组合货品 API请求
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
type AlibabadchainaoxiangcombinescitembatchupdateAPIRequest struct {
	model.Params
	// 批量更新组合货品入参
	_batchUpdateCombineScitemRequest *BatchUpdateCombineScItemRequest
}

// NewAlibabadchainaoxiangcombinescitembatchupdateRequest 初始化AlibabadchainaoxiangcombinescitembatchupdateAPIRequest对象
func NewAlibabadchainaoxiangcombinescitembatchupdateRequest() *AlibabadchainaoxiangcombinescitembatchupdateAPIRequest {
	return &AlibabadchainaoxiangcombinescitembatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangcombinescitembatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.combinescitem.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangcombinescitembatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangcombinescitembatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchUpdateCombineScitemRequest is BatchUpdateCombineScitemRequest Setter
// 批量更新组合货品入参
func (r *AlibabadchainaoxiangcombinescitembatchupdateAPIRequest) SetBatchUpdateCombineScitemRequest(_batchUpdateCombineScitemRequest *BatchUpdateCombineScItemRequest) error {
	r._batchUpdateCombineScitemRequest = _batchUpdateCombineScitemRequest
	r.Set("batch_update_combine_scitem_request", _batchUpdateCombineScitemRequest)
	return nil
}

// GetBatchUpdateCombineScitemRequest BatchUpdateCombineScitemRequest Getter
func (r AlibabadchainaoxiangcombinescitembatchupdateAPIRequest) GetBatchUpdateCombineScitemRequest() *BatchUpdateCombineScItemRequest {
	return r._batchUpdateCombineScitemRequest
}
