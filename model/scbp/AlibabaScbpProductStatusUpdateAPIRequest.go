package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpproductstatusupdateAPIRequest 修改P4P产品推广状态 API请求
// alibaba.scbp.product.status.update
//
// 修改P4P产品推广状态
type AlibabascbpproductstatusupdateAPIRequest struct {
	model.Params
	// 产品ID列表
	_productIdList []string
	// enabled:开启,disabled:暂停
	_status string
}

// NewAlibabascbpproductstatusupdateRequest 初始化AlibabascbpproductstatusupdateAPIRequest对象
func NewAlibabascbpproductstatusupdateRequest() *AlibabascbpproductstatusupdateAPIRequest {
	return &AlibabascbpproductstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpproductstatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpproductstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpproductstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIdList is ProductIdList Setter
// 产品ID列表
func (r *AlibabascbpproductstatusupdateAPIRequest) SetProductIdList(_productIdList []string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// GetProductIdList ProductIdList Getter
func (r AlibabascbpproductstatusupdateAPIRequest) GetProductIdList() []string {
	return r._productIdList
}

// SetStatus is Status Setter
// enabled:开启,disabled:暂停
func (r *AlibabascbpproductstatusupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabascbpproductstatusupdateAPIRequest) GetStatus() string {
	return r._status
}
