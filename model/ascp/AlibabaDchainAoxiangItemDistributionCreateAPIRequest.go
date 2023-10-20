package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributioncreateAPIRequest 选择店铺商品并进行铺货（通用） API请求
// alibaba.dchain.aoxiang.item.distribution.create
//
// 选择店铺商品并进行铺货, 铺货给所有的合作分销商。设定的价格为通用价格
type AlibabadchainaoxiangitemdistributioncreateAPIRequest struct {
	model.Params
	// 创建商品分销入参
	_createItemDistributionRequest *CreateItemDistributionRequest
}

// NewAlibabadchainaoxiangitemdistributioncreateRequest 初始化AlibabadchainaoxiangitemdistributioncreateAPIRequest对象
func NewAlibabadchainaoxiangitemdistributioncreateRequest() *AlibabadchainaoxiangitemdistributioncreateAPIRequest {
	return &AlibabadchainaoxiangitemdistributioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemdistributioncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemdistributioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemdistributioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateItemDistributionRequest is CreateItemDistributionRequest Setter
// 创建商品分销入参
func (r *AlibabadchainaoxiangitemdistributioncreateAPIRequest) SetCreateItemDistributionRequest(_createItemDistributionRequest *CreateItemDistributionRequest) error {
	r._createItemDistributionRequest = _createItemDistributionRequest
	r.Set("create_item_distribution_request", _createItemDistributionRequest)
	return nil
}

// GetCreateItemDistributionRequest CreateItemDistributionRequest Getter
func (r AlibabadchainaoxiangitemdistributioncreateAPIRequest) GetCreateItemDistributionRequest() *CreateItemDistributionRequest {
	return r._createItemDistributionRequest
}
