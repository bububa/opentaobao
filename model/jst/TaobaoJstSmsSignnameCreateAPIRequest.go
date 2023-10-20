package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameCreateAPIRequest 淘宝短信签名创建接口 API请求
// taobao.jst.sms.signname.create
//
// 聚石塔短信签名创建接口
type TaobaoJstSmsSignnameCreateAPIRequest struct {
	model.Params
	// 创建签名入参
	_addSmsSignRequest *TopAddSmsSignRequest
}

// NewTaobaoJstSmsSignnameCreateRequest 初始化TaobaoJstSmsSignnameCreateAPIRequest对象
func NewTaobaoJstSmsSignnameCreateRequest() *TaobaoJstSmsSignnameCreateAPIRequest {
	return &TaobaoJstSmsSignnameCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsSignnameCreateAPIRequest) Reset() {
	r._addSmsSignRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsSignnameCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddSmsSignRequest is AddSmsSignRequest Setter
// 创建签名入参
func (r *TaobaoJstSmsSignnameCreateAPIRequest) SetAddSmsSignRequest(_addSmsSignRequest *TopAddSmsSignRequest) error {
	r._addSmsSignRequest = _addSmsSignRequest
	r.Set("add_sms_sign_request", _addSmsSignRequest)
	return nil
}

// GetAddSmsSignRequest AddSmsSignRequest Getter
func (r TaobaoJstSmsSignnameCreateAPIRequest) GetAddSmsSignRequest() *TopAddSmsSignRequest {
	return r._addSmsSignRequest
}

var poolTaobaoJstSmsSignnameCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsSignnameCreateRequest()
	},
}

// GetTaobaoJstSmsSignnameCreateRequest 从 sync.Pool 获取 TaobaoJstSmsSignnameCreateAPIRequest
func GetTaobaoJstSmsSignnameCreateAPIRequest() *TaobaoJstSmsSignnameCreateAPIRequest {
	return poolTaobaoJstSmsSignnameCreateAPIRequest.Get().(*TaobaoJstSmsSignnameCreateAPIRequest)
}

// ReleaseTaobaoJstSmsSignnameCreateAPIRequest 将 TaobaoJstSmsSignnameCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsSignnameCreateAPIRequest(v *TaobaoJstSmsSignnameCreateAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsSignnameCreateAPIRequest.Put(v)
}
