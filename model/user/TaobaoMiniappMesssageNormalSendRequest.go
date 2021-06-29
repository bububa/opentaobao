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
type TaobaoMiniappMesssageNormalSendRequest struct {
    model.Params
    // 普通消息结构
    _param   *DownNormalMessageDto
}

// 初始化TaobaoMiniappMesssageNormalSendRequest对象
func NewTaobaoMiniappMesssageNormalSendRequest() *TaobaoMiniappMesssageNormalSendRequest{
    return &TaobaoMiniappMesssageNormalSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappMesssageNormalSendRequest) GetApiMethodName() string {
    return "taobao.miniapp.messsage.normal.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappMesssageNormalSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 普通消息结构
func (r *TaobaoMiniappMesssageNormalSendRequest) SetParam(_param *DownNormalMessageDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoMiniappMesssageNormalSendRequest) GetParam() *DownNormalMessageDto {
    return r._param
}
