package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitembatchupdateAPIRequest alibaba.dchain.aoxiang.scitem.batch.update API请求
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
type AlibabadchainaoxiangscitembatchupdateAPIRequest struct {
	model.Params
	// 批量更新货品入参，最多30条
	_batchUpdateScitemRequest *BatchUpdateScItemRequest
}

// NewAlibabadchainaoxiangscitembatchupdateRequest 初始化AlibabadchainaoxiangscitembatchupdateAPIRequest对象
func NewAlibabadchainaoxiangscitembatchupdateRequest() *AlibabadchainaoxiangscitembatchupdateAPIRequest {
	return &AlibabadchainaoxiangscitembatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangscitembatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangscitembatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangscitembatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchUpdateScitemRequest is BatchUpdateScitemRequest Setter
// 批量更新货品入参，最多30条
func (r *AlibabadchainaoxiangscitembatchupdateAPIRequest) SetBatchUpdateScitemRequest(_batchUpdateScitemRequest *BatchUpdateScItemRequest) error {
	r._batchUpdateScitemRequest = _batchUpdateScitemRequest
	r.Set("batch_update_scitem_request", _batchUpdateScitemRequest)
	return nil
}

// GetBatchUpdateScitemRequest BatchUpdateScitemRequest Getter
func (r AlibabadchainaoxiangscitembatchupdateAPIRequest) GetBatchUpdateScitemRequest() *BatchUpdateScItemRequest {
	return r._batchUpdateScitemRequest
}
