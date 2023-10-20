package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionDescribeAPIRequest 分销商品文描 API请求
// alibaba.dchain.aoxiang.item.distribution.describe
//
// 分销商品文描
type AlibabaDchainAoxiangItemDistributionDescribeAPIRequest struct {
	model.Params
	// 分销文描入参
	_createItemDistributionRequest *MaterialRequest
}

// NewAlibabaDchainAoxiangItemDistributionDescribeRequest 初始化AlibabaDchainAoxiangItemDistributionDescribeAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionDescribeRequest() *AlibabaDchainAoxiangItemDistributionDescribeAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionDescribeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) Reset() {
	r._createItemDistributionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.describe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateItemDistributionRequest is CreateItemDistributionRequest Setter
// 分销文描入参
func (r *AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) SetCreateItemDistributionRequest(_createItemDistributionRequest *MaterialRequest) error {
	r._createItemDistributionRequest = _createItemDistributionRequest
	r.Set("create_item_distribution_request", _createItemDistributionRequest)
	return nil
}

// GetCreateItemDistributionRequest CreateItemDistributionRequest Getter
func (r AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) GetCreateItemDistributionRequest() *MaterialRequest {
	return r._createItemDistributionRequest
}

var poolAlibabaDchainAoxiangItemDistributionDescribeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemDistributionDescribeRequest()
	},
}

// GetAlibabaDchainAoxiangItemDistributionDescribeRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionDescribeAPIRequest
func GetAlibabaDchainAoxiangItemDistributionDescribeAPIRequest() *AlibabaDchainAoxiangItemDistributionDescribeAPIRequest {
	return poolAlibabaDchainAoxiangItemDistributionDescribeAPIRequest.Get().(*AlibabaDchainAoxiangItemDistributionDescribeAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemDistributionDescribeAPIRequest 将 AlibabaDchainAoxiangItemDistributionDescribeAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionDescribeAPIRequest(v *AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionDescribeAPIRequest.Put(v)
}
