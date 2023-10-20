package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurprcreateAPIRequest 下pr单 API请求
// alibaba.pur.pr.create
//
// 下pr单
type AlibabapurprcreateAPIRequest struct {
	model.Params
	// 订单信息
	_purReq *MallReceivePrRequest
}

// NewAlibabapurprcreateRequest 初始化AlibabapurprcreateAPIRequest对象
func NewAlibabapurprcreateRequest() *AlibabapurprcreateAPIRequest {
	return &AlibabapurprcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurprcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.pr.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurprcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurprcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPurReq is PurReq Setter
// 订单信息
func (r *AlibabapurprcreateAPIRequest) SetPurReq(_purReq *MallReceivePrRequest) error {
	r._purReq = _purReq
	r.Set("pur_req", _purReq)
	return nil
}

// GetPurReq PurReq Getter
func (r AlibabapurprcreateAPIRequest) GetPurReq() *MallReceivePrRequest {
	return r._purReq
}
