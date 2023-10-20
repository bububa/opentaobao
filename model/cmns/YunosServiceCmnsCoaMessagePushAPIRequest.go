package cmns

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessagePushAPIRequest 消息推送接口 API请求
// yunos.service.cmns.coa.message.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
type YunosServiceCmnsCoaMessagePushAPIRequest struct {
	model.Params
	// 消息推送请求对象
	_pushRequest *PushRequest
}

// NewYunosServiceCmnsCoaMessagePushRequest 初始化YunosServiceCmnsCoaMessagePushAPIRequest对象
func NewYunosServiceCmnsCoaMessagePushRequest() *YunosServiceCmnsCoaMessagePushAPIRequest {
	return &YunosServiceCmnsCoaMessagePushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaMessagePushAPIRequest) Reset() {
	r._pushRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushRequest is PushRequest Setter
// 消息推送请求对象
func (r *YunosServiceCmnsCoaMessagePushAPIRequest) SetPushRequest(_pushRequest *PushRequest) error {
	r._pushRequest = _pushRequest
	r.Set("push_request", _pushRequest)
	return nil
}

// GetPushRequest PushRequest Getter
func (r YunosServiceCmnsCoaMessagePushAPIRequest) GetPushRequest() *PushRequest {
	return r._pushRequest
}

var poolYunosServiceCmnsCoaMessagePushAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaMessagePushRequest()
	},
}

// GetYunosServiceCmnsCoaMessagePushRequest 从 sync.Pool 获取 YunosServiceCmnsCoaMessagePushAPIRequest
func GetYunosServiceCmnsCoaMessagePushAPIRequest() *YunosServiceCmnsCoaMessagePushAPIRequest {
	return poolYunosServiceCmnsCoaMessagePushAPIRequest.Get().(*YunosServiceCmnsCoaMessagePushAPIRequest)
}

// ReleaseYunosServiceCmnsCoaMessagePushAPIRequest 将 YunosServiceCmnsCoaMessagePushAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaMessagePushAPIRequest(v *YunosServiceCmnsCoaMessagePushAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaMessagePushAPIRequest.Put(v)
}
