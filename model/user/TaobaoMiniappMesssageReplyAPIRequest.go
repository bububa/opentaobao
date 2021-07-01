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
type TaobaoMiniappMesssageReplyAPIRequest struct {
    model.Params
    // 入参
    _param   *ReplyMessageDto
}

// 初始化TaobaoMiniappMesssageReplyAPIRequest对象
func NewTaobaoMiniappMesssageReplyRequest() *TaobaoMiniappMesssageReplyAPIRequest{
    return &TaobaoMiniappMesssageReplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappMesssageReplyAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.messsage.reply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappMesssageReplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *TaobaoMiniappMesssageReplyAPIRequest) SetParam(_param *ReplyMessageDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoMiniappMesssageReplyAPIRequest) GetParam() *ReplyMessageDto {
    return r._param
}
