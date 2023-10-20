package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoamessagepushAPIRequest 消息推送接口 API请求
// yunos.service.cmns.coa.message.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
type YunosservicecmnscoamessagepushAPIRequest struct {
	model.Params
	// 消息推送请求对象
	_pushRequest *PushRequest
}

// NewYunosservicecmnscoamessagepushRequest 初始化YunosservicecmnscoamessagepushAPIRequest对象
func NewYunosservicecmnscoamessagepushRequest() *YunosservicecmnscoamessagepushAPIRequest {
	return &YunosservicecmnscoamessagepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoamessagepushAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoamessagepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoamessagepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushRequest is PushRequest Setter
// 消息推送请求对象
func (r *YunosservicecmnscoamessagepushAPIRequest) SetPushRequest(_pushRequest *PushRequest) error {
	r._pushRequest = _pushRequest
	r.Set("push_request", _pushRequest)
	return nil
}

// GetPushRequest PushRequest Getter
func (r YunosservicecmnscoamessagepushAPIRequest) GetPushRequest() *PushRequest {
	return r._pushRequest
}
