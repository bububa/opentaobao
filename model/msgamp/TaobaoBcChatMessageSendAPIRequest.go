package msgamp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBcChatMessageSendAPIRequest 小程序资源授权-BC客服消息 API请求
// taobao.bc.chat.message.send
//
// 小程序资源授权-消息订阅
type TaobaoBcChatMessageSendAPIRequest struct {
	model.Params
	// 资源授权token
	_topRecourceToken string
	// 资源点
	_topRecourceId string
	// 请求参数
	_msgRequest *MiniappBcChatMsgRequest
}

// NewTaobaoBcChatMessageSendRequest 初始化TaobaoBcChatMessageSendAPIRequest对象
func NewTaobaoBcChatMessageSendRequest() *TaobaoBcChatMessageSendAPIRequest {
	return &TaobaoBcChatMessageSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBcChatMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.bc.chat.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBcChatMessageSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBcChatMessageSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopRecourceToken is TopRecourceToken Setter
// 资源授权token
func (r *TaobaoBcChatMessageSendAPIRequest) SetTopRecourceToken(_topRecourceToken string) error {
	r._topRecourceToken = _topRecourceToken
	r.Set("top_recource_token", _topRecourceToken)
	return nil
}

// GetTopRecourceToken TopRecourceToken Getter
func (r TaobaoBcChatMessageSendAPIRequest) GetTopRecourceToken() string {
	return r._topRecourceToken
}

// SetTopRecourceId is TopRecourceId Setter
// 资源点
func (r *TaobaoBcChatMessageSendAPIRequest) SetTopRecourceId(_topRecourceId string) error {
	r._topRecourceId = _topRecourceId
	r.Set("top_recource_id", _topRecourceId)
	return nil
}

// GetTopRecourceId TopRecourceId Getter
func (r TaobaoBcChatMessageSendAPIRequest) GetTopRecourceId() string {
	return r._topRecourceId
}

// SetMsgRequest is MsgRequest Setter
// 请求参数
func (r *TaobaoBcChatMessageSendAPIRequest) SetMsgRequest(_msgRequest *MiniappBcChatMsgRequest) error {
	r._msgRequest = _msgRequest
	r.Set("msg_request", _msgRequest)
	return nil
}

// GetMsgRequest MsgRequest Getter
func (r TaobaoBcChatMessageSendAPIRequest) GetMsgRequest() *MiniappBcChatMsgRequest {
	return r._msgRequest
}
