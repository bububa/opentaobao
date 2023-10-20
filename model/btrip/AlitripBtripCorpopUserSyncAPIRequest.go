package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopusersyncAPIRequest 外部人员同步 API请求
// alitrip.btrip.corpop.user.sync
//
// 同步外部平台用户信息至商旅内部
type AlitripbtripcorpopusersyncAPIRequest struct {
	model.Params
	// 人员同步请求
	_rq *BtripUserSyncRq
}

// NewAlitripbtripcorpopusersyncRequest 初始化AlitripbtripcorpopusersyncAPIRequest对象
func NewAlitripbtripcorpopusersyncRequest() *AlitripbtripcorpopusersyncAPIRequest {
	return &AlitripbtripcorpopusersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopusersyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopusersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopusersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 人员同步请求
func (r *AlitripbtripcorpopusersyncAPIRequest) SetRq(_rq *BtripUserSyncRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopusersyncAPIRequest) GetRq() *BtripUserSyncRq {
	return r._rq
}
