package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemBatchCreateAPIRequest 新建货品 API请求
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
type AlibabaDchainAoxiangScitemBatchCreateAPIRequest struct {
	model.Params
	// 批量新建货品入参，数量不大于30
	_batchCreateScitemRequest *BatchCreateScItemRequest
}

// NewAlibabaDchainAoxiangScitemBatchCreateRequest 初始化AlibabaDchainAoxiangScitemBatchCreateAPIRequest对象
func NewAlibabaDchainAoxiangScitemBatchCreateRequest() *AlibabaDchainAoxiangScitemBatchCreateAPIRequest {
	return &AlibabaDchainAoxiangScitemBatchCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemBatchCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.batch.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemBatchCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangScitemBatchCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchCreateScitemRequest is BatchCreateScitemRequest Setter
// 批量新建货品入参，数量不大于30
func (r *AlibabaDchainAoxiangScitemBatchCreateAPIRequest) SetBatchCreateScitemRequest(_batchCreateScitemRequest *BatchCreateScItemRequest) error {
	r._batchCreateScitemRequest = _batchCreateScitemRequest
	r.Set("batch_create_scitem_request", _batchCreateScitemRequest)
	return nil
}

// GetBatchCreateScitemRequest BatchCreateScitemRequest Getter
func (r AlibabaDchainAoxiangScitemBatchCreateAPIRequest) GetBatchCreateScitemRequest() *BatchCreateScItemRequest {
	return r._batchCreateScitemRequest
}
