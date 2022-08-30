package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemBatchUpdateAPIRequest alibaba.dchain.aoxiang.scitem.batch.update API请求
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
type AlibabaDchainAoxiangScitemBatchUpdateAPIRequest struct {
	model.Params
	// 批量更新货品入参，最多30条
	_batchUpdateScitemRequest *BatchUpdateScItemRequest
}

// NewAlibabaDchainAoxiangScitemBatchUpdateRequest 初始化AlibabaDchainAoxiangScitemBatchUpdateAPIRequest对象
func NewAlibabaDchainAoxiangScitemBatchUpdateRequest() *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest {
	return &AlibabaDchainAoxiangScitemBatchUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBatchUpdateScitemRequest is BatchUpdateScitemRequest Setter
// 批量更新货品入参，最多30条
func (r *AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) SetBatchUpdateScitemRequest(_batchUpdateScitemRequest *BatchUpdateScItemRequest) error {
	r._batchUpdateScitemRequest = _batchUpdateScitemRequest
	r.Set("batch_update_scitem_request", _batchUpdateScitemRequest)
	return nil
}

// GetBatchUpdateScitemRequest BatchUpdateScitemRequest Getter
func (r AlibabaDchainAoxiangScitemBatchUpdateAPIRequest) GetBatchUpdateScitemRequest() *BatchUpdateScItemRequest {
	return r._batchUpdateScitemRequest
}
