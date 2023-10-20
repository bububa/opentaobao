package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoapushAPIRequest 消息推送接口 API请求
// yunos.service.cmns.coa.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
type YunosservicecmnscoapushAPIRequest struct {
	model.Params
	// 消息结构对象
	_msgObj *CmnsMessage
}

// NewYunosservicecmnscoapushRequest 初始化YunosservicecmnscoapushAPIRequest对象
func NewYunosservicecmnscoapushRequest() *YunosservicecmnscoapushAPIRequest {
	return &YunosservicecmnscoapushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoapushAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoapushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoapushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsgObj is MsgObj Setter
// 消息结构对象
func (r *YunosservicecmnscoapushAPIRequest) SetMsgObj(_msgObj *CmnsMessage) error {
	r._msgObj = _msgObj
	r.Set("msg_obj", _msgObj)
	return nil
}

// GetMsgObj MsgObj Getter
func (r YunosservicecmnscoapushAPIRequest) GetMsgObj() *CmnsMessage {
	return r._msgObj
}
