package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest 小游戏发送push API请求
// alibaba.alsc.growth.interactive.mini.game.notice.push.send
//
// 小游戏发送push
type AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest struct {
	model.Params
	// push发送请求参数
	_pushSendRequest *PushSendRequest
}

// NewAlibabaAlscGrowthInteractiveMiniGameNoticePushSendRequest 初始化AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest对象
func NewAlibabaAlscGrowthInteractiveMiniGameNoticePushSendRequest() *AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest {
	return &AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.mini.game.notice.push.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushSendRequest is PushSendRequest Setter
// push发送请求参数
func (r *AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest) SetPushSendRequest(_pushSendRequest *PushSendRequest) error {
	r._pushSendRequest = _pushSendRequest
	r.Set("push_send_request", _pushSendRequest)
	return nil
}

// GetPushSendRequest PushSendRequest Getter
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest) GetPushSendRequest() *PushSendRequest {
	return r._pushSendRequest
}
