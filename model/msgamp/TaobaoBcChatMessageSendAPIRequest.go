package msgamp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobcchatmessagesendAPIRequest 小程序资源授权-BC客服消息 API请求
// taobao.bc.chat.message.send
//
// 小程序资源授权-消息订阅
type TaobaobcchatmessagesendAPIRequest struct {
	model.Params
	// 资源授权token
	_topRecourceToken string
	// 资源点
	_topRecourceId string
	// 请求参数
	_msgRequest *MiniappBcChatMsgRequest
}

// NewTaobaobcchatmessagesendRequest 初始化TaobaobcchatmessagesendAPIRequest对象
func NewTaobaobcchatmessagesendRequest() *TaobaobcchatmessagesendAPIRequest {
	return &TaobaobcchatmessagesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobcchatmessagesendAPIRequest) GetApiMethodName() string {
	return "taobao.bc.chat.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobcchatmessagesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobcchatmessagesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopRecourceToken is TopRecourceToken Setter
// 资源授权token
func (r *TaobaobcchatmessagesendAPIRequest) SetTopRecourceToken(_topRecourceToken string) error {
	r._topRecourceToken = _topRecourceToken
	r.Set("top_recource_token", _topRecourceToken)
	return nil
}

// GetTopRecourceToken TopRecourceToken Getter
func (r TaobaobcchatmessagesendAPIRequest) GetTopRecourceToken() string {
	return r._topRecourceToken
}

// SetTopRecourceId is TopRecourceId Setter
// 资源点
func (r *TaobaobcchatmessagesendAPIRequest) SetTopRecourceId(_topRecourceId string) error {
	r._topRecourceId = _topRecourceId
	r.Set("top_recource_id", _topRecourceId)
	return nil
}

// GetTopRecourceId TopRecourceId Getter
func (r TaobaobcchatmessagesendAPIRequest) GetTopRecourceId() string {
	return r._topRecourceId
}

// SetMsgRequest is MsgRequest Setter
// 请求参数
func (r *TaobaobcchatmessagesendAPIRequest) SetMsgRequest(_msgRequest *MiniappBcChatMsgRequest) error {
	r._msgRequest = _msgRequest
	r.Set("msg_request", _msgRequest)
	return nil
}

// GetMsgRequest MsgRequest Getter
func (r TaobaobcchatmessagesendAPIRequest) GetMsgRequest() *MiniappBcChatMsgRequest {
	return r._msgRequest
}
