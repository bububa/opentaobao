package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 APIRequest
yunos.service.cmns.coa.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaPushRequest struct {
    model.Params

    // 消息结构对象
    msgObj   *CmnsMessage 

}

func NewYunosServiceCmnsCoaPushRequest() *YunosServiceCmnsCoaPushRequest{
    return &YunosServiceCmnsCoaPushRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaPushRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.push"
}

func (r YunosServiceCmnsCoaPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaPushRequest) SetMsgObj(msgObj *CmnsMessage) error {
    r.msgObj = msgObj
    r.Set("msg_obj", msgObj)
    return nil
}

func (r YunosServiceCmnsCoaPushRequest) GetMsgObj() *CmnsMessage {
    return r.msgObj
}

