package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopUserSyncAPIRequest
外部人员同步 API请求
alitrip.btrip.corpop.user.sync

同步外部平台用户信息至商旅内部 */
type AlitripBtripCorpopUserSyncAPIRequest struct {
	model.Params
	// 人员同步请求
	_rq *BtripUserSyncRq
}

// NewAlitripBtripCorpopUserSyncRequest 初始化AlitripBtripCorpopUserSyncAPIRequest对象
func NewAlitripBtripCorpopUserSyncRequest() *AlitripBtripCorpopUserSyncAPIRequest {
	return &AlitripBtripCorpopUserSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopUserSyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopUserSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 人员同步请求
func (r *AlitripBtripCorpopUserSyncAPIRequest) SetRq(_rq *BtripUserSyncRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCorpopUserSyncAPIRequest) GetRq() *BtripUserSyncRq {
	return r._rq
}
