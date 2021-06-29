package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除创意 API请求
taobao.simba.creative.delete

删除一个创意
*/
type TaobaoSimbaCreativeDeleteRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 创意Id
    creativeId   int64
}

// 初始化TaobaoSimbaCreativeDeleteRequest对象
func NewTaobaoSimbaCreativeDeleteRequest() *TaobaoSimbaCreativeDeleteRequest{
    return &TaobaoSimbaCreativeDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.creative.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativeDeleteRequest) GetNick() string {
    return r.nick
}
// CreativeId Setter
// 创意Id
func (r *TaobaoSimbaCreativeDeleteRequest) SetCreativeId(creativeId int64) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoSimbaCreativeDeleteRequest) GetCreativeId() int64 {
    return r.creativeId
}
