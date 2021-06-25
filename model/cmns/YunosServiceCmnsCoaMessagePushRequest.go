package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 APIRequest
yunos.service.cmns.coa.message.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaMessagePushRequest struct {
    model.Params

    // 消息推送请求对象
    pushRequest   *PushRequest 

}

func NewYunosServiceCmnsCoaMessagePushRequest() *YunosServiceCmnsCoaMessagePushRequest{
    return &YunosServiceCmnsCoaMessagePushRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaMessagePushRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.push"
}

func (r YunosServiceCmnsCoaMessagePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaMessagePushRequest) SetPushRequest(pushRequest *PushRequest) error {
    r.pushRequest = pushRequest
    r.Set("push_request", pushRequest)
    return nil
}

func (r YunosServiceCmnsCoaMessagePushRequest) GetPushRequest() *PushRequest {
    return r.pushRequest
}

