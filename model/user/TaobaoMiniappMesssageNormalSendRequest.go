package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行普通消息给用户 API请求
taobao.miniapp.messsage.normal.send

小程序下行单个普通消息
*/
type TaobaoMiniappMesssageNormalSendAPIRequest struct {
    model.Params
    // 普通消息结构
    _param   *DownNormalMessageDTO
}

// 初始化TaobaoMiniappMesssageNormalSendAPIRequest对象
func NewTaobaoMiniappMesssageNormalSendRequest() *TaobaoMiniappMesssageNormalSendAPIRequest{
    return &TaobaoMiniappMesssageNormalSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.messsage.normal.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 普通消息结构
func (r *TaobaoMiniappMesssageNormalSendAPIRequest) SetParam(_param *DownNormalMessageDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoMiniappMesssageNormalSendAPIRequest) GetParam() *DownNormalMessageDTO {
    return r._param
}
