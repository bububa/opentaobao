package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaPushAPIRequest
消息推送接口 API请求
yunos.service.cmns.coa.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。 */
type YunosServiceCmnsCoaPushAPIRequest struct {
	model.Params
	// 消息结构对象
	_msgObj *CmnsMessage
}

// NewYunosServiceCmnsCoaPushRequest 初始化YunosServiceCmnsCoaPushAPIRequest对象
func NewYunosServiceCmnsCoaPushRequest() *YunosServiceCmnsCoaPushAPIRequest {
	return &YunosServiceCmnsCoaPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaPushAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MsgObj Setter
// 消息结构对象
func (r *YunosServiceCmnsCoaPushAPIRequest) SetMsgObj(_msgObj *CmnsMessage) error {
	r._msgObj = _msgObj
	r.Set("msg_obj", _msgObj)
	return nil
}

// Get MsgObj Getter
func (r YunosServiceCmnsCoaPushAPIRequest) GetMsgObj() *CmnsMessage {
	return r._msgObj
}
