package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送 API请求
tmall.fans.arena.push

超级擂台消息推送
*/
type TmallFansArenaPushRequest struct {
    model.Params
    // 推送列表
    _pushList   []PushMessageParamDo
}

// 初始化TmallFansArenaPushRequest对象
func NewTmallFansArenaPushRequest() *TmallFansArenaPushRequest{
    return &TmallFansArenaPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansArenaPushRequest) GetApiMethodName() string {
    return "tmall.fans.arena.push"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansArenaPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushList Setter
// 推送列表
func (r *TmallFansArenaPushRequest) SetPushList(_pushList []PushMessageParamDo) error {
    r._pushList = _pushList
    r.Set("push_list", _pushList)
    return nil
}

// PushList Getter
func (r TmallFansArenaPushRequest) GetPushList() []PushMessageParamDo {
    return r._pushList
}
