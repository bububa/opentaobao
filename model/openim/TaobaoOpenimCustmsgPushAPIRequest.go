package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimCustmsgPushAPIRequest 推送自定义openim消息 API请求
// taobao.openim.custmsg.push
//
// isv通过该接口给openim用户推送自定义消息
type TaobaoOpenimCustmsgPushAPIRequest struct {
	model.Params
	// 自定义消息内容
	_custmsg *CustMsg
}

// NewTaobaoOpenimCustmsgPushRequest 初始化TaobaoOpenimCustmsgPushAPIRequest对象
func NewTaobaoOpenimCustmsgPushRequest() *TaobaoOpenimCustmsgPushAPIRequest {
	return &TaobaoOpenimCustmsgPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimCustmsgPushAPIRequest) GetApiMethodName() string {
	return "taobao.openim.custmsg.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimCustmsgPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Custmsg Setter
// 自定义消息内容
func (r *TaobaoOpenimCustmsgPushAPIRequest) SetCustmsg(_custmsg *CustMsg) error {
	r._custmsg = _custmsg
	r.Set("custmsg", _custmsg)
	return nil
}

// Get Custmsg Getter
func (r TaobaoOpenimCustmsgPushAPIRequest) GetCustmsg() *CustMsg {
	return r._custmsg
}
