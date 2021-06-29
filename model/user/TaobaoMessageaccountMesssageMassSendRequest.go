package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息号开放-消息群发 API请求
taobao.messageaccount.messsage.mass.send

外部 isv 调用该进口来进行消息号消息的群发
*/
type TaobaoMessageaccountMesssageMassSendRequest struct {
    model.Params
    // 参数
    param   *MassMessageDto
}

// 初始化TaobaoMessageaccountMesssageMassSendRequest对象
func NewTaobaoMessageaccountMesssageMassSendRequest() *TaobaoMessageaccountMesssageMassSendRequest{
    return &TaobaoMessageaccountMesssageMassSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageMassSendRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.mass.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageMassSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *TaobaoMessageaccountMesssageMassSendRequest) SetParam(param *MassMessageDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r TaobaoMessageaccountMesssageMassSendRequest) GetParam() *MassMessageDto {
    return r.param
}
