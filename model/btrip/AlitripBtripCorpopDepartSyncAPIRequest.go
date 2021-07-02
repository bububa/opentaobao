package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopDepartSyncAPIRequest 外部部门同步 API请求
// alitrip.btrip.corpop.depart.sync
//
// 同步外部平台部门信息至商旅内部
type AlitripBtripCorpopDepartSyncAPIRequest struct {
	model.Params
	// 同步部门请求
	_rq *BtripDepartSyncRq
}

// NewAlitripBtripCorpopDepartSyncRequest 初始化AlitripBtripCorpopDepartSyncAPIRequest对象
func NewAlitripBtripCorpopDepartSyncRequest() *AlitripBtripCorpopDepartSyncAPIRequest {
	return &AlitripBtripCorpopDepartSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopDepartSyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.depart.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopDepartSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 同步部门请求
func (r *AlitripBtripCorpopDepartSyncAPIRequest) SetRq(_rq *BtripDepartSyncRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopDepartSyncAPIRequest) GetRq() *BtripDepartSyncRq {
	return r._rq
}
