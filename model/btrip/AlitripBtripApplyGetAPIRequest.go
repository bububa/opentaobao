package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApplyGetAPIRequest 获取单个审批单 API请求
// alitrip.btrip.apply.get
//
// 获取单个审批单的详情数据
type AlitripBtripApplyGetAPIRequest struct {
	model.Params
	// 外部审批单id
	_thirdpartApplyId string
	// 企业id
	_corpId string
	// 审批单展示id
	_applyShowId string
	// 阿里商旅审批单id
	_applyId int64
}

// NewAlitripBtripApplyGetRequest 初始化AlitripBtripApplyGetAPIRequest对象
func NewAlitripBtripApplyGetRequest() *AlitripBtripApplyGetAPIRequest {
	return &AlitripBtripApplyGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripApplyGetAPIRequest) Reset() {
	r._thirdpartApplyId = ""
	r._corpId = ""
	r._applyShowId = ""
	r._applyId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripApplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripApplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripApplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdpartApplyId is ThirdpartApplyId Setter
// 外部审批单id
func (r *AlitripBtripApplyGetAPIRequest) SetThirdpartApplyId(_thirdpartApplyId string) error {
	r._thirdpartApplyId = _thirdpartApplyId
	r.Set("thirdpart_apply_id", _thirdpartApplyId)
	return nil
}

// GetThirdpartApplyId ThirdpartApplyId Getter
func (r AlitripBtripApplyGetAPIRequest) GetThirdpartApplyId() string {
	return r._thirdpartApplyId
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripBtripApplyGetAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripBtripApplyGetAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetApplyShowId is ApplyShowId Setter
// 审批单展示id
func (r *AlitripBtripApplyGetAPIRequest) SetApplyShowId(_applyShowId string) error {
	r._applyShowId = _applyShowId
	r.Set("apply_show_id", _applyShowId)
	return nil
}

// GetApplyShowId ApplyShowId Getter
func (r AlitripBtripApplyGetAPIRequest) GetApplyShowId() string {
	return r._applyShowId
}

// SetApplyId is ApplyId Setter
// 阿里商旅审批单id
func (r *AlitripBtripApplyGetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripBtripApplyGetAPIRequest) GetApplyId() int64 {
	return r._applyId
}

var poolAlitripBtripApplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripApplyGetRequest()
	},
}

// GetAlitripBtripApplyGetRequest 从 sync.Pool 获取 AlitripBtripApplyGetAPIRequest
func GetAlitripBtripApplyGetAPIRequest() *AlitripBtripApplyGetAPIRequest {
	return poolAlitripBtripApplyGetAPIRequest.Get().(*AlitripBtripApplyGetAPIRequest)
}

// ReleaseAlitripBtripApplyGetAPIRequest 将 AlitripBtripApplyGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripApplyGetAPIRequest(v *AlitripBtripApplyGetAPIRequest) {
	v.Reset()
	poolAlitripBtripApplyGetAPIRequest.Put(v)
}
