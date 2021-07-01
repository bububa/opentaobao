package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimImmsgPushAPIRequest
openim标准消息发送 API请求
taobao.openim.immsg.push

服务端对openim用户发送标准消息，包括文字、语音、图片等。 */
type TaobaoOpenimImmsgPushAPIRequest struct {
	model.Params
	// openim消息结构体
	_immsg *ImMsg
}

// NewTaobaoOpenimImmsgPushRequest 初始化TaobaoOpenimImmsgPushAPIRequest对象
func NewTaobaoOpenimImmsgPushRequest() *TaobaoOpenimImmsgPushAPIRequest {
	return &TaobaoOpenimImmsgPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimImmsgPushAPIRequest) GetApiMethodName() string {
	return "taobao.openim.immsg.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimImmsgPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Immsg Setter
// openim消息结构体
func (r *TaobaoOpenimImmsgPushAPIRequest) SetImmsg(_immsg *ImMsg) error {
	r._immsg = _immsg
	r.Set("immsg", _immsg)
	return nil
}

// Get Immsg Getter
func (r TaobaoOpenimImmsgPushAPIRequest) GetImmsg() *ImMsg {
	return r._immsg
}
