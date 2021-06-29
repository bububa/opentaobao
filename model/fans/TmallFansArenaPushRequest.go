package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送 APIRequest
tmall.fans.arena.push

超级擂台消息推送
*/
type TmallFansArenaPushRequest struct {
    model.Params

    // 推送列表
    pushList   []PushMessageParamDo 

}

func NewTmallFansArenaPushRequest() *TmallFansArenaPushRequest{
    return &TmallFansArenaPushRequest{
        Params: model.NewParams(),
    }
}

func (r TmallFansArenaPushRequest) GetApiMethodName() string {
    return "tmall.fans.arena.push"
}

func (r TmallFansArenaPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallFansArenaPushRequest) SetPushList(pushList []PushMessageParamDo) error {
    r.pushList = pushList
    r.Set("push_list", pushList)
    return nil
}

func (r TmallFansArenaPushRequest) GetPushList() []PushMessageParamDo {
    return r.pushList
}

