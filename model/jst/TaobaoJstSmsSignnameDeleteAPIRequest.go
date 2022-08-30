package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameDeleteAPIRequest 淘宝短信签名删除 API请求
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
type TaobaoJstSmsSignnameDeleteAPIRequest struct {
	model.Params
	// 删除签名入参
	_deleteSmsSignRequest *TopDeleteSmsSignRequest
}

// NewTaobaoJstSmsSignnameDeleteRequest 初始化TaobaoJstSmsSignnameDeleteAPIRequest对象
func NewTaobaoJstSmsSignnameDeleteRequest() *TaobaoJstSmsSignnameDeleteAPIRequest {
	return &TaobaoJstSmsSignnameDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeleteSmsSignRequest is DeleteSmsSignRequest Setter
// 删除签名入参
func (r *TaobaoJstSmsSignnameDeleteAPIRequest) SetDeleteSmsSignRequest(_deleteSmsSignRequest *TopDeleteSmsSignRequest) error {
	r._deleteSmsSignRequest = _deleteSmsSignRequest
	r.Set("delete_sms_sign_request", _deleteSmsSignRequest)
	return nil
}

// GetDeleteSmsSignRequest DeleteSmsSignRequest Getter
func (r TaobaoJstSmsSignnameDeleteAPIRequest) GetDeleteSmsSignRequest() *TopDeleteSmsSignRequest {
	return r._deleteSmsSignRequest
}
