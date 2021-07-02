package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurPrCreateAPIRequest 下pr单 API请求
// alibaba.pur.pr.create
//
// 下pr单
type AlibabaPurPrCreateAPIRequest struct {
	model.Params
	// 订单信息
	_purReq *MallReceivePrRequest
}

// NewAlibabaPurPrCreateRequest 初始化AlibabaPurPrCreateAPIRequest对象
func NewAlibabaPurPrCreateRequest() *AlibabaPurPrCreateAPIRequest {
	return &AlibabaPurPrCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurPrCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.pr.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurPrCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPurReq is PurReq Setter
// 订单信息
func (r *AlibabaPurPrCreateAPIRequest) SetPurReq(_purReq *MallReceivePrRequest) error {
	r._purReq = _purReq
	r.Set("pur_req", _purReq)
	return nil
}

// GetPurReq PurReq Getter
func (r AlibabaPurPrCreateAPIRequest) GetPurReq() *MallReceivePrRequest {
	return r._purReq
}
