package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest 取消商品分销 API请求
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
type AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest struct {
	model.Params
	// 取消商品分销入参
	_cancelDistributeRequest *CancelDistributeRequest
}

// NewAlibabadchainaoxiangitemdistributionbatchcancelRequest 初始化AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest对象
func NewAlibabadchainaoxiangitemdistributionbatchcancelRequest() *AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest {
	return &AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.batch.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelDistributeRequest is CancelDistributeRequest Setter
// 取消商品分销入参
func (r *AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest) SetCancelDistributeRequest(_cancelDistributeRequest *CancelDistributeRequest) error {
	r._cancelDistributeRequest = _cancelDistributeRequest
	r.Set("cancel_distribute_request", _cancelDistributeRequest)
	return nil
}

// GetCancelDistributeRequest CancelDistributeRequest Getter
func (r AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest) GetCancelDistributeRequest() *CancelDistributeRequest {
	return r._cancelDistributeRequest
}
