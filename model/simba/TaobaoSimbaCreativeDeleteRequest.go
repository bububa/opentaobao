package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除创意 APIRequest
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

func NewTaobaoSimbaCreativeDeleteRequest() *TaobaoSimbaCreativeDeleteRequest{
    return &TaobaoSimbaCreativeDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCreativeDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.creative.delete"
}

func (r TaobaoSimbaCreativeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCreativeDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCreativeDeleteRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCreativeDeleteRequest) SetCreativeId(creativeId int64) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r TaobaoSimbaCreativeDeleteRequest) GetCreativeId() int64 {
    return r.creativeId
}

