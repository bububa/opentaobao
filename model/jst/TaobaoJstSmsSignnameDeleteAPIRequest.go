package jst

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsSignnameDeleteAPIRequest) Reset() {
	r._deleteSmsSignRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsSignnameDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoJstSmsSignnameDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsSignnameDeleteRequest()
	},
}

// GetTaobaoJstSmsSignnameDeleteRequest 从 sync.Pool 获取 TaobaoJstSmsSignnameDeleteAPIRequest
func GetTaobaoJstSmsSignnameDeleteAPIRequest() *TaobaoJstSmsSignnameDeleteAPIRequest {
	return poolTaobaoJstSmsSignnameDeleteAPIRequest.Get().(*TaobaoJstSmsSignnameDeleteAPIRequest)
}

// ReleaseTaobaoJstSmsSignnameDeleteAPIRequest 将 TaobaoJstSmsSignnameDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsSignnameDeleteAPIRequest(v *TaobaoJstSmsSignnameDeleteAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsSignnameDeleteAPIRequest.Put(v)
}
