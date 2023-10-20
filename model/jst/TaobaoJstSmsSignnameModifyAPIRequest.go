package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameModifyAPIRequest 淘宝短信签名修改 API请求
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
type TaobaoJstSmsSignnameModifyAPIRequest struct {
	model.Params
	// 修改签名入参
	_modifySmsSignRequest *TopModifySmsSignRequest
}

// NewTaobaoJstSmsSignnameModifyRequest 初始化TaobaoJstSmsSignnameModifyAPIRequest对象
func NewTaobaoJstSmsSignnameModifyRequest() *TaobaoJstSmsSignnameModifyAPIRequest {
	return &TaobaoJstSmsSignnameModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsSignnameModifyAPIRequest) Reset() {
	r._modifySmsSignRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifySmsSignRequest is ModifySmsSignRequest Setter
// 修改签名入参
func (r *TaobaoJstSmsSignnameModifyAPIRequest) SetModifySmsSignRequest(_modifySmsSignRequest *TopModifySmsSignRequest) error {
	r._modifySmsSignRequest = _modifySmsSignRequest
	r.Set("modify_sms_sign_request", _modifySmsSignRequest)
	return nil
}

// GetModifySmsSignRequest ModifySmsSignRequest Getter
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetModifySmsSignRequest() *TopModifySmsSignRequest {
	return r._modifySmsSignRequest
}

var poolTaobaoJstSmsSignnameModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsSignnameModifyRequest()
	},
}

// GetTaobaoJstSmsSignnameModifyRequest 从 sync.Pool 获取 TaobaoJstSmsSignnameModifyAPIRequest
func GetTaobaoJstSmsSignnameModifyAPIRequest() *TaobaoJstSmsSignnameModifyAPIRequest {
	return poolTaobaoJstSmsSignnameModifyAPIRequest.Get().(*TaobaoJstSmsSignnameModifyAPIRequest)
}

// ReleaseTaobaoJstSmsSignnameModifyAPIRequest 将 TaobaoJstSmsSignnameModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsSignnameModifyAPIRequest(v *TaobaoJstSmsSignnameModifyAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsSignnameModifyAPIRequest.Put(v)
}
