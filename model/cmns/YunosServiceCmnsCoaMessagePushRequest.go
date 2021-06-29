package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 API请求
yunos.service.cmns.coa.message.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaMessagePushRequest struct {
    model.Params
    // 消息推送请求对象
    _pushRequest   *PushRequest
}

// 初始化YunosServiceCmnsCoaMessagePushRequest对象
func NewYunosServiceCmnsCoaMessagePushRequest() *YunosServiceCmnsCoaMessagePushRequest{
    return &YunosServiceCmnsCoaMessagePushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessagePushRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.push"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessagePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushRequest Setter
// 消息推送请求对象
func (r *YunosServiceCmnsCoaMessagePushRequest) SetPushRequest(_pushRequest *PushRequest) error {
    r._pushRequest = _pushRequest
    r.Set("push_request", _pushRequest)
    return nil
}

// PushRequest Getter
func (r YunosServiceCmnsCoaMessagePushRequest) GetPushRequest() *PushRequest {
    return r._pushRequest
}
