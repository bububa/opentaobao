package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCallRecordAPIRequest 回访记录回传API API请求
// tmall.servicecenter.workcard.call.record
//
// 客满回访信息登记
type TmallServicecenterWorkcardCallRecordAPIRequest struct {
	model.Params
	// 请求入参
	_busiRequest *UpdateAttributeRequest
}

// NewTmallServicecenterWorkcardCallRecordRequest 初始化TmallServicecenterWorkcardCallRecordAPIRequest对象
func NewTmallServicecenterWorkcardCallRecordRequest() *TmallServicecenterWorkcardCallRecordAPIRequest {
	return &TmallServicecenterWorkcardCallRecordAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardCallRecordAPIRequest) Reset() {
	r._busiRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.call.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusiRequest is BusiRequest Setter
// 请求入参
func (r *TmallServicecenterWorkcardCallRecordAPIRequest) SetBusiRequest(_busiRequest *UpdateAttributeRequest) error {
	r._busiRequest = _busiRequest
	r.Set("busi_request", _busiRequest)
	return nil
}

// GetBusiRequest BusiRequest Getter
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetBusiRequest() *UpdateAttributeRequest {
	return r._busiRequest
}

var poolTmallServicecenterWorkcardCallRecordAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardCallRecordRequest()
	},
}

// GetTmallServicecenterWorkcardCallRecordRequest 从 sync.Pool 获取 TmallServicecenterWorkcardCallRecordAPIRequest
func GetTmallServicecenterWorkcardCallRecordAPIRequest() *TmallServicecenterWorkcardCallRecordAPIRequest {
	return poolTmallServicecenterWorkcardCallRecordAPIRequest.Get().(*TmallServicecenterWorkcardCallRecordAPIRequest)
}

// ReleaseTmallServicecenterWorkcardCallRecordAPIRequest 将 TmallServicecenterWorkcardCallRecordAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardCallRecordAPIRequest(v *TmallServicecenterWorkcardCallRecordAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardCallRecordAPIRequest.Put(v)
}
