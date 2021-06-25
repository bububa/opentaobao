package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息号开放-消息群发 APIRequest
taobao.messageaccount.messsage.mass.send

外部 isv 调用该进口来进行消息号消息的群发
*/
type TaobaoMessageaccountMesssageMassSendRequest struct {
    model.Params

    // 参数
    param   *MassMessageDto 

}

func NewTaobaoMessageaccountMesssageMassSendRequest() *TaobaoMessageaccountMesssageMassSendRequest{
    return &TaobaoMessageaccountMesssageMassSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMessageaccountMesssageMassSendRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.mass.send"
}

func (r TaobaoMessageaccountMesssageMassSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMessageaccountMesssageMassSendRequest) SetParam(param *MassMessageDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoMessageaccountMesssageMassSendRequest) GetParam() *MassMessageDto {
    return r.param
}

