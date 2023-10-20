package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributiondescribeAPIRequest 分销商品文描 API请求
// alibaba.dchain.aoxiang.item.distribution.describe
//
// 分销商品文描
type AlibabadchainaoxiangitemdistributiondescribeAPIRequest struct {
	model.Params
	// 分销文描入参
	_createItemDistributionRequest *MaterialRequest
}

// NewAlibabadchainaoxiangitemdistributiondescribeRequest 初始化AlibabadchainaoxiangitemdistributiondescribeAPIRequest对象
func NewAlibabadchainaoxiangitemdistributiondescribeRequest() *AlibabadchainaoxiangitemdistributiondescribeAPIRequest {
	return &AlibabadchainaoxiangitemdistributiondescribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemdistributiondescribeAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.describe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemdistributiondescribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemdistributiondescribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateItemDistributionRequest is CreateItemDistributionRequest Setter
// 分销文描入参
func (r *AlibabadchainaoxiangitemdistributiondescribeAPIRequest) SetCreateItemDistributionRequest(_createItemDistributionRequest *MaterialRequest) error {
	r._createItemDistributionRequest = _createItemDistributionRequest
	r.Set("create_item_distribution_request", _createItemDistributionRequest)
	return nil
}

// GetCreateItemDistributionRequest CreateItemDistributionRequest Getter
func (r AlibabadchainaoxiangitemdistributiondescribeAPIRequest) GetCreateItemDistributionRequest() *MaterialRequest {
	return r._createItemDistributionRequest
}
