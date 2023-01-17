package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest 批量发送push API请求
// alibaba.alsc.growth.interactive.mini.game.notice.push.batch.send
//
// 批量发送push
type AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest struct {
	model.Params
	// 批量发送消息请求对象
	_pushSendBatchRequest *PushSendBatchRequest
}

// NewAlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendRequest 初始化AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest对象
func NewAlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendRequest() *AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest {
	return &AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.mini.game.notice.push.batch.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushSendBatchRequest is PushSendBatchRequest Setter
// 批量发送消息请求对象
func (r *AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest) SetPushSendBatchRequest(_pushSendBatchRequest *PushSendBatchRequest) error {
	r._pushSendBatchRequest = _pushSendBatchRequest
	r.Set("push_send_batch_request", _pushSendBatchRequest)
	return nil
}

// GetPushSendBatchRequest PushSendBatchRequest Getter
func (r AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest) GetPushSendBatchRequest() *PushSendBatchRequest {
	return r._pushSendBatchRequest
}
