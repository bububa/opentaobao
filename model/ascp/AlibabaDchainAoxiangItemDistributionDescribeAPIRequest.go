package ascp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.describe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionDescribeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
