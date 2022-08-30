package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardTpconfirmAPIRequest 确认服务完成 API请求
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
type TmallServicecenterWorkcardTpconfirmAPIRequest struct {
	model.Params
	// 服务商确认服务完成信息
	_tpConfirmRequest *TpConfirmRequest
}

// NewTmallServicecenterWorkcardTpconfirmRequest 初始化TmallServicecenterWorkcardTpconfirmAPIRequest对象
func NewTmallServicecenterWorkcardTpconfirmRequest() *TmallServicecenterWorkcardTpconfirmAPIRequest {
	return &TmallServicecenterWorkcardTpconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.tpconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTpConfirmRequest is TpConfirmRequest Setter
// 服务商确认服务完成信息
func (r *TmallServicecenterWorkcardTpconfirmAPIRequest) SetTpConfirmRequest(_tpConfirmRequest *TpConfirmRequest) error {
	r._tpConfirmRequest = _tpConfirmRequest
	r.Set("tp_confirm_request", _tpConfirmRequest)
	return nil
}

// GetTpConfirmRequest TpConfirmRequest Getter
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetTpConfirmRequest() *TpConfirmRequest {
	return r._tpConfirmRequest
}
