package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopExceedapplySyncAPIRequest 第三方超标审批结果回传 API请求
// alitrip.btrip.corpop.exceedapply.sync
//
// 第三方审批单推送到企业后，企业审批结束，将审批结果回传给阿里商旅
type AlitripBtripCorpopExceedapplySyncAPIRequest struct {
	model.Params
	// 入参
	_rq *BtripExceedApplyRq
}

// NewAlitripBtripCorpopExceedapplySyncRequest 初始化AlitripBtripCorpopExceedapplySyncAPIRequest对象
func NewAlitripBtripCorpopExceedapplySyncRequest() *AlitripBtripCorpopExceedapplySyncAPIRequest {
	return &AlitripBtripCorpopExceedapplySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.exceedapply.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopExceedapplySyncAPIRequest) SetRq(_rq *BtripExceedApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopExceedapplySyncAPIRequest) GetRq() *BtripExceedApplyRq {
	return r._rq
}
