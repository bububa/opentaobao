package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopUserSyncAPIRequest 外部人员同步 API请求
// alitrip.btrip.corpop.user.sync
//
// 同步外部平台用户信息至商旅内部
type AlitripBtripCorpopUserSyncAPIRequest struct {
	model.Params
	// 人员同步请求
	_rq *BtripUserSyncRq
}

// NewAlitripBtripCorpopUserSyncRequest 初始化AlitripBtripCorpopUserSyncAPIRequest对象
func NewAlitripBtripCorpopUserSyncRequest() *AlitripBtripCorpopUserSyncAPIRequest {
	return &AlitripBtripCorpopUserSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopUserSyncAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopUserSyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopUserSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopUserSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 人员同步请求
func (r *AlitripBtripCorpopUserSyncAPIRequest) SetRq(_rq *BtripUserSyncRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopUserSyncAPIRequest) GetRq() *BtripUserSyncRq {
	return r._rq
}

var poolAlitripBtripCorpopUserSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopUserSyncRequest()
	},
}

// GetAlitripBtripCorpopUserSyncRequest 从 sync.Pool 获取 AlitripBtripCorpopUserSyncAPIRequest
func GetAlitripBtripCorpopUserSyncAPIRequest() *AlitripBtripCorpopUserSyncAPIRequest {
	return poolAlitripBtripCorpopUserSyncAPIRequest.Get().(*AlitripBtripCorpopUserSyncAPIRequest)
}

// ReleaseAlitripBtripCorpopUserSyncAPIRequest 将 AlitripBtripCorpopUserSyncAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopUserSyncAPIRequest(v *AlitripBtripCorpopUserSyncAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopUserSyncAPIRequest.Put(v)
}
