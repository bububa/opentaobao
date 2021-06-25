package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行普通消息给用户 APIRequest
taobao.miniapp.messsage.normal.send

小程序下行单个普通消息
*/
type TaobaoMiniappMesssageNormalSendRequest struct {
    model.Params

    // 普通消息结构
    param   *DownNormalMessageDto 

}

func NewTaobaoMiniappMesssageNormalSendRequest() *TaobaoMiniappMesssageNormalSendRequest{
    return &TaobaoMiniappMesssageNormalSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappMesssageNormalSendRequest) GetApiMethodName() string {
    return "taobao.miniapp.messsage.normal.send"
}

func (r TaobaoMiniappMesssageNormalSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappMesssageNormalSendRequest) SetParam(param *DownNormalMessageDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoMiniappMesssageNormalSendRequest) GetParam() *DownNormalMessageDto {
    return r.param
}

