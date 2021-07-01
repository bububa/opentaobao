package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessagePushAPIRequest
消息推送接口 API请求
yunos.service.cmns.coa.message.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。 */
type YunosServiceCmnsCoaMessagePushAPIRequest struct {
	model.Params
	// 消息推送请求对象
	_pushRequest *PushRequest
}

// NewYunosServiceCmnsCoaMessagePushRequest 初始化YunosServiceCmnsCoaMessagePushAPIRequest对象
func NewYunosServiceCmnsCoaMessagePushRequest() *YunosServiceCmnsCoaMessagePushAPIRequest {
	return &YunosServiceCmnsCoaMessagePushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushRequest Setter
// 消息推送请求对象
func (r *YunosServiceCmnsCoaMessagePushAPIRequest) SetPushRequest(_pushRequest *PushRequest) error {
	r._pushRequest = _pushRequest
	r.Set("push_request", _pushRequest)
	return nil
}

// Get PushRequest Getter
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetPushRequest() *PushRequest {
	return r._pushRequest
}
