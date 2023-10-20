package cmns

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaPushAPIRequest 消息推送接口 API请求
// yunos.service.cmns.coa.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
type YunosServiceCmnsCoaPushAPIRequest struct {
	model.Params
	// 消息结构对象
	_msgObj *CmnsMessage
}

// NewYunosServiceCmnsCoaPushRequest 初始化YunosServiceCmnsCoaPushAPIRequest对象
func NewYunosServiceCmnsCoaPushRequest() *YunosServiceCmnsCoaPushAPIRequest {
	return &YunosServiceCmnsCoaPushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaPushAPIRequest) Reset() {
	r._msgObj = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaPushAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsgObj is MsgObj Setter
// 消息结构对象
func (r *YunosServiceCmnsCoaPushAPIRequest) SetMsgObj(_msgObj *CmnsMessage) error {
	r._msgObj = _msgObj
	r.Set("msg_obj", _msgObj)
	return nil
}

// GetMsgObj MsgObj Getter
func (r YunosServiceCmnsCoaPushAPIRequest) GetMsgObj() *CmnsMessage {
	return r._msgObj
}

var poolYunosServiceCmnsCoaPushAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaPushRequest()
	},
}

// GetYunosServiceCmnsCoaPushRequest 从 sync.Pool 获取 YunosServiceCmnsCoaPushAPIRequest
func GetYunosServiceCmnsCoaPushAPIRequest() *YunosServiceCmnsCoaPushAPIRequest {
	return poolYunosServiceCmnsCoaPushAPIRequest.Get().(*YunosServiceCmnsCoaPushAPIRequest)
}

// ReleaseYunosServiceCmnsCoaPushAPIRequest 将 YunosServiceCmnsCoaPushAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaPushAPIRequest(v *YunosServiceCmnsCoaPushAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaPushAPIRequest.Put(v)
}
