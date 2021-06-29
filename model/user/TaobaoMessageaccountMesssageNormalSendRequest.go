package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下行普通消息 API请求
taobao.messageaccount.messsage.normal.send

消息号下行单个普通消息
*/
type TaobaoMessageaccountMesssageNormalSendRequest struct {
    model.Params
    // 下行普通消息
    _param   *NormalMessageDTO
}

// 初始化TaobaoMessageaccountMesssageNormalSendRequest对象
func NewTaobaoMessageaccountMesssageNormalSendRequest() *TaobaoMessageaccountMesssageNormalSendRequest{
    return &TaobaoMessageaccountMesssageNormalSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMessageaccountMesssageNormalSendRequest) GetApiMethodName() string {
    return "taobao.messageaccount.messsage.normal.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMessageaccountMesssageNormalSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 下行普通消息
func (r *TaobaoMessageaccountMesssageNormalSendRequest) SetParam(_param *NormalMessageDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoMessageaccountMesssageNormalSendRequest) GetParam() *NormalMessageDTO {
    return r._param
}
