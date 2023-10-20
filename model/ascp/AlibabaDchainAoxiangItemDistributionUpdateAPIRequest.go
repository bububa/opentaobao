package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributionupdateAPIRequest 更新商品分销内容 API请求
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
type AlibabadchainaoxiangitemdistributionupdateAPIRequest struct {
	model.Params
	// 更新商品分销入参
	_updateItemDistributionRequest *UpdateItemDistributionRequest
}

// NewAlibabadchainaoxiangitemdistributionupdateRequest 初始化AlibabadchainaoxiangitemdistributionupdateAPIRequest对象
func NewAlibabadchainaoxiangitemdistributionupdateRequest() *AlibabadchainaoxiangitemdistributionupdateAPIRequest {
	return &AlibabadchainaoxiangitemdistributionupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemdistributionupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemdistributionupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemdistributionupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateItemDistributionRequest is UpdateItemDistributionRequest Setter
// 更新商品分销入参
func (r *AlibabadchainaoxiangitemdistributionupdateAPIRequest) SetUpdateItemDistributionRequest(_updateItemDistributionRequest *UpdateItemDistributionRequest) error {
	r._updateItemDistributionRequest = _updateItemDistributionRequest
	r.Set("update_item_distribution_request", _updateItemDistributionRequest)
	return nil
}

// GetUpdateItemDistributionRequest UpdateItemDistributionRequest Getter
func (r AlibabadchainaoxiangitemdistributionupdateAPIRequest) GetUpdateItemDistributionRequest() *UpdateItemDistributionRequest {
	return r._updateItemDistributionRequest
}
