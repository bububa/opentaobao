package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardtpconfirmAPIRequest 确认服务完成 API请求
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
type TmallservicecenterworkcardtpconfirmAPIRequest struct {
	model.Params
	// 服务商确认服务完成信息
	_tpConfirmRequest *TpConfirmRequest
}

// NewTmallservicecenterworkcardtpconfirmRequest 初始化TmallservicecenterworkcardtpconfirmAPIRequest对象
func NewTmallservicecenterworkcardtpconfirmRequest() *TmallservicecenterworkcardtpconfirmAPIRequest {
	return &TmallservicecenterworkcardtpconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardtpconfirmAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.tpconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardtpconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardtpconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTpConfirmRequest is TpConfirmRequest Setter
// 服务商确认服务完成信息
func (r *TmallservicecenterworkcardtpconfirmAPIRequest) SetTpConfirmRequest(_tpConfirmRequest *TpConfirmRequest) error {
	r._tpConfirmRequest = _tpConfirmRequest
	r.Set("tp_confirm_request", _tpConfirmRequest)
	return nil
}

// GetTpConfirmRequest TpConfirmRequest Getter
func (r TmallservicecenterworkcardtpconfirmAPIRequest) GetTpConfirmRequest() *TpConfirmRequest {
	return r._tpConfirmRequest
}
