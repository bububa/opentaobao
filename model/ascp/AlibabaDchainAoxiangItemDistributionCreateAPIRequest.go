package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionCreateAPIRequest 选择店铺商品并进行铺货（通用） API请求
// alibaba.dchain.aoxiang.item.distribution.create
//
// 选择店铺商品并进行铺货, 铺货给所有的合作分销商。设定的价格为通用价格
type AlibabaDchainAoxiangItemDistributionCreateAPIRequest struct {
	model.Params
	// 创建商品分销入参
	_createItemDistributionRequest *CreateItemDistributionRequest
}

// NewAlibabaDchainAoxiangItemDistributionCreateRequest 初始化AlibabaDchainAoxiangItemDistributionCreateAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionCreateRequest() *AlibabaDchainAoxiangItemDistributionCreateAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemDistributionCreateAPIRequest) Reset() {
	r._createItemDistributionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemDistributionCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateItemDistributionRequest is CreateItemDistributionRequest Setter
// 创建商品分销入参
func (r *AlibabaDchainAoxiangItemDistributionCreateAPIRequest) SetCreateItemDistributionRequest(_createItemDistributionRequest *CreateItemDistributionRequest) error {
	r._createItemDistributionRequest = _createItemDistributionRequest
	r.Set("create_item_distribution_request", _createItemDistributionRequest)
	return nil
}

// GetCreateItemDistributionRequest CreateItemDistributionRequest Getter
func (r AlibabaDchainAoxiangItemDistributionCreateAPIRequest) GetCreateItemDistributionRequest() *CreateItemDistributionRequest {
	return r._createItemDistributionRequest
}

var poolAlibabaDchainAoxiangItemDistributionCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemDistributionCreateRequest()
	},
}

// GetAlibabaDchainAoxiangItemDistributionCreateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionCreateAPIRequest
func GetAlibabaDchainAoxiangItemDistributionCreateAPIRequest() *AlibabaDchainAoxiangItemDistributionCreateAPIRequest {
	return poolAlibabaDchainAoxiangItemDistributionCreateAPIRequest.Get().(*AlibabaDchainAoxiangItemDistributionCreateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemDistributionCreateAPIRequest 将 AlibabaDchainAoxiangItemDistributionCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionCreateAPIRequest(v *AlibabaDchainAoxiangItemDistributionCreateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionCreateAPIRequest.Put(v)
}
