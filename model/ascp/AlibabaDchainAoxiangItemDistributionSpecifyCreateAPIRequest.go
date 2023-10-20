package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest 指定分销商进行铺货(专享) API请求
// alibaba.dchain.aoxiang.item.distribution.specify.create
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
type AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest struct {
	model.Params
	// 指定分销商铺品
	_createItemDistributionRequest *SpecifyDistributionRequest
}

// NewAlibabaDchainAoxiangItemDistributionSpecifyCreateRequest 初始化AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionSpecifyCreateRequest() *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) Reset() {
	r._createItemDistributionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.specify.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateItemDistributionRequest is CreateItemDistributionRequest Setter
// 指定分销商铺品
func (r *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) SetCreateItemDistributionRequest(_createItemDistributionRequest *SpecifyDistributionRequest) error {
	r._createItemDistributionRequest = _createItemDistributionRequest
	r.Set("create_item_distribution_request", _createItemDistributionRequest)
	return nil
}

// GetCreateItemDistributionRequest CreateItemDistributionRequest Getter
func (r AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) GetCreateItemDistributionRequest() *SpecifyDistributionRequest {
	return r._createItemDistributionRequest
}

var poolAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemDistributionSpecifyCreateRequest()
	},
}

// GetAlibabaDchainAoxiangItemDistributionSpecifyCreateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest
func GetAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest() *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest {
	return poolAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest.Get().(*AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest 将 AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest(v *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIRequest.Put(v)
}
