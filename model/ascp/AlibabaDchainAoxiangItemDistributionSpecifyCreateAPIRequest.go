package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest 指定分销商进行铺货(专享) API请求
// alibaba.dchain.aoxiang.item.distribution.specify.create
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
type AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest struct {
	model.Params
	// 指定分销商铺品
	_createItemDistributionRequest *SpecifyDistributionRequest
}

// NewAlibabadchainaoxiangitemdistributionspecifycreateRequest 初始化AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest对象
func NewAlibabadchainaoxiangitemdistributionspecifycreateRequest() *AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest {
	return &AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.specify.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateItemDistributionRequest is CreateItemDistributionRequest Setter
// 指定分销商铺品
func (r *AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest) SetCreateItemDistributionRequest(_createItemDistributionRequest *SpecifyDistributionRequest) error {
	r._createItemDistributionRequest = _createItemDistributionRequest
	r.Set("create_item_distribution_request", _createItemDistributionRequest)
	return nil
}

// GetCreateItemDistributionRequest CreateItemDistributionRequest Getter
func (r AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest) GetCreateItemDistributionRequest() *SpecifyDistributionRequest {
	return r._createItemDistributionRequest
}
