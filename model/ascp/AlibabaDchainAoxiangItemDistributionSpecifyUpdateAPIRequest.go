package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest 指定分销商进行铺货(专享) - 修改 API请求
// alibaba.dchain.aoxiang.item.distribution.specify.update
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
type AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest struct {
	model.Params
	// 创建商品分销入参
	_updateItemDistributionRequest *UpdateItemDistributionRequest
}

// NewAlibabaDchainAoxiangItemDistributionSpecifyUpdateRequest 初始化AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionSpecifyUpdateRequest() *AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) Reset() {
	r._updateItemDistributionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.specify.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateItemDistributionRequest is UpdateItemDistributionRequest Setter
// 创建商品分销入参
func (r *AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) SetUpdateItemDistributionRequest(_updateItemDistributionRequest *UpdateItemDistributionRequest) error {
	r._updateItemDistributionRequest = _updateItemDistributionRequest
	r.Set("update_item_distribution_request", _updateItemDistributionRequest)
	return nil
}

// GetUpdateItemDistributionRequest UpdateItemDistributionRequest Getter
func (r AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) GetUpdateItemDistributionRequest() *UpdateItemDistributionRequest {
	return r._updateItemDistributionRequest
}

var poolAlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemDistributionSpecifyUpdateRequest()
	},
}

// GetAlibabaDchainAoxiangItemDistributionSpecifyUpdateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest
func GetAlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest() *AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest {
	return poolAlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest.Get().(*AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest 将 AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest(v *AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest.Put(v)
}
