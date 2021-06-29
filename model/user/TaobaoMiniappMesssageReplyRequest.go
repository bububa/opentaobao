package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻店铺下行回复接口 API请求
taobao.miniapp.messsage.reply

外部 isv 调用该进口来进行轻店铺消息的回复
*/
type TaobaoMiniappMesssageReplyRequest struct {
    model.Params
    // 入参
    param   *ReplyMessageDto
}

// 初始化TaobaoMiniappMesssageReplyRequest对象
func NewTaobaoMiniappMesssageReplyRequest() *TaobaoMiniappMesssageReplyRequest{
    return &TaobaoMiniappMesssageReplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappMesssageReplyRequest) GetApiMethodName() string {
    return "taobao.miniapp.messsage.reply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappMesssageReplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *TaobaoMiniappMesssageReplyRequest) SetParam(param *ReplyMessageDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r TaobaoMiniappMesssageReplyRequest) GetParam() *ReplyMessageDto {
    return r.param
}
