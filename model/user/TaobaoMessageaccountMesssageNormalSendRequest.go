package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下行普通消息 APIRequest
taobao.messageaccount.messsage.normal.send

消息号下行单个普通消息
*/
type TaobaoMessageaccountMesssageNormalSendRequest struct {
    model.Params

    // 下行普通消息
    param   *NormalMessageDto 

}

func NewTaobaoMessageaccountMesssageNormalSendRequest() *TaobaoMessageaccountMesssageNormalSendRequest{
    return &TaobaoMessageaccountMesssageNormalSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMessageaccountMesssageNormalSendRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.normal.send"
}

func (r TaobaoMessageaccountMesssageNormalSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMessageaccountMesssageNormalSendRequest) SetParam(param *NormalMessageDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoMessageaccountMesssageNormalSendRequest) GetParam() *NormalMessageDto {
    return r.param
}

