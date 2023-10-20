package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopDepartSyncAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopDepartSyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.depart.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopDepartSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopDepartSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripCorpopDepartSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopDepartSyncRequest()
	},
}

// GetAlitripBtripCorpopDepartSyncRequest 从 sync.Pool 获取 AlitripBtripCorpopDepartSyncAPIRequest
func GetAlitripBtripCorpopDepartSyncAPIRequest() *AlitripBtripCorpopDepartSyncAPIRequest {
	return poolAlitripBtripCorpopDepartSyncAPIRequest.Get().(*AlitripBtripCorpopDepartSyncAPIRequest)
}

// ReleaseAlitripBtripCorpopDepartSyncAPIRequest 将 AlitripBtripCorpopDepartSyncAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopDepartSyncAPIRequest(v *AlitripBtripCorpopDepartSyncAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopDepartSyncAPIRequest.Put(v)
}
