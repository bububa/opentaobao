package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitembatchcreateAPIRequest 新建货品 API请求
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
type AlibabadchainaoxiangscitembatchcreateAPIRequest struct {
	model.Params
	// 批量新建货品入参，数量不大于30
	_batchCreateScitemRequest *BatchCreateScItemRequest
}

// NewAlibabadchainaoxiangscitembatchcreateRequest 初始化AlibabadchainaoxiangscitembatchcreateAPIRequest对象
func NewAlibabadchainaoxiangscitembatchcreateRequest() *AlibabadchainaoxiangscitembatchcreateAPIRequest {
	return &AlibabadchainaoxiangscitembatchcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangscitembatchcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.batch.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangscitembatchcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangscitembatchcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchCreateScitemRequest is BatchCreateScitemRequest Setter
// 批量新建货品入参，数量不大于30
func (r *AlibabadchainaoxiangscitembatchcreateAPIRequest) SetBatchCreateScitemRequest(_batchCreateScitemRequest *BatchCreateScItemRequest) error {
	r._batchCreateScitemRequest = _batchCreateScitemRequest
	r.Set("batch_create_scitem_request", _batchCreateScitemRequest)
	return nil
}

// GetBatchCreateScitemRequest BatchCreateScitemRequest Getter
func (r AlibabadchainaoxiangscitembatchcreateAPIRequest) GetBatchCreateScitemRequest() *BatchCreateScItemRequest {
	return r._batchCreateScitemRequest
}
