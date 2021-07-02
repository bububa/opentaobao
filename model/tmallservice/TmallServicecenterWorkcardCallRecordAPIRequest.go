package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.call.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BusiRequest Setter
// 请求入参
func (r *TmallServicecenterWorkcardCallRecordAPIRequest) SetBusiRequest(_busiRequest *UpdateAttributeRequest) error {
	r._busiRequest = _busiRequest
	r.Set("busi_request", _busiRequest)
	return nil
}

// Get BusiRequest Getter
func (r TmallServicecenterWorkcardCallRecordAPIRequest) GetBusiRequest() *UpdateAttributeRequest {
	return r._busiRequest
}
