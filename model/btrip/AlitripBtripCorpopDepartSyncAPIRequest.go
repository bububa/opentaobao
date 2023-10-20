package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopdepartsyncAPIRequest 外部部门同步 API请求
// alitrip.btrip.corpop.depart.sync
//
// 同步外部平台部门信息至商旅内部
type AlitripbtripcorpopdepartsyncAPIRequest struct {
	model.Params
	// 同步部门请求
	_rq *BtripDepartSyncRq
}

// NewAlitripbtripcorpopdepartsyncRequest 初始化AlitripbtripcorpopdepartsyncAPIRequest对象
func NewAlitripbtripcorpopdepartsyncRequest() *AlitripbtripcorpopdepartsyncAPIRequest {
	return &AlitripbtripcorpopdepartsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopdepartsyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.depart.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopdepartsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopdepartsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 同步部门请求
func (r *AlitripbtripcorpopdepartsyncAPIRequest) SetRq(_rq *BtripDepartSyncRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopdepartsyncAPIRequest) GetRq() *BtripDepartSyncRq {
	return r._rq
}
