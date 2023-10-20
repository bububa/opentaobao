package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest 取消商品分销 API请求
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
type AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest struct {
	model.Params
	// 取消商品分销入参
	_cancelDistributeRequest *CancelDistributeRequest
}

// NewAlibabaDchainAoxiangItemDistributionBatchCancelRequest 初始化AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionBatchCancelRequest() *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) Reset() {
	r._cancelDistributeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.batch.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelDistributeRequest is CancelDistributeRequest Setter
// 取消商品分销入参
func (r *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) SetCancelDistributeRequest(_cancelDistributeRequest *CancelDistributeRequest) error {
	r._cancelDistributeRequest = _cancelDistributeRequest
	r.Set("cancel_distribute_request", _cancelDistributeRequest)
	return nil
}

// GetCancelDistributeRequest CancelDistributeRequest Getter
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetCancelDistributeRequest() *CancelDistributeRequest {
	return r._cancelDistributeRequest
}

var poolAlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemDistributionBatchCancelRequest()
	},
}

// GetAlibabaDchainAoxiangItemDistributionBatchCancelRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest
func GetAlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest() *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest {
	return poolAlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest.Get().(*AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest 将 AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest(v *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest.Put(v)
}
