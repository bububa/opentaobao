package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardcallrecordAPIRequest 回访记录回传API API请求
// tmall.servicecenter.workcard.call.record
//
// 客满回访信息登记
type TmallservicecenterworkcardcallrecordAPIRequest struct {
	model.Params
	// 请求入参
	_busiRequest *UpdateAttributeRequest
}

// NewTmallservicecenterworkcardcallrecordRequest 初始化TmallservicecenterworkcardcallrecordAPIRequest对象
func NewTmallservicecenterworkcardcallrecordRequest() *TmallservicecenterworkcardcallrecordAPIRequest {
	return &TmallservicecenterworkcardcallrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardcallrecordAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.call.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardcallrecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardcallrecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusiRequest is BusiRequest Setter
// 请求入参
func (r *TmallservicecenterworkcardcallrecordAPIRequest) SetBusiRequest(_busiRequest *UpdateAttributeRequest) error {
	r._busiRequest = _busiRequest
	r.Set("busi_request", _busiRequest)
	return nil
}

// GetBusiRequest BusiRequest Getter
func (r TmallservicecenterworkcardcallrecordAPIRequest) GetBusiRequest() *UpdateAttributeRequest {
	return r._busiRequest
}
