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
type TmallFansArenaPushAPIRequest struct {
    model.Params
    // 推送列表
    _pushList   []PushMessageParamDO
}

// 初始化TmallFansArenaPushAPIRequest对象
func NewTmallFansArenaPushRequest() *TmallFansArenaPushAPIRequest{
    return &TmallFansArenaPushAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansArenaPushAPIRequest) GetApiMethodName() string {
    return "tmall.fans.arena.push"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansArenaPushAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushList Setter
// 推送列表
func (r *TmallFansArenaPushAPIRequest) SetPushList(_pushList []PushMessageParamDO) error {
    r._pushList = _pushList
    r.Set("push_list", _pushList)
    return nil
}

// PushList Getter
func (r TmallFansArenaPushAPIRequest) GetPushList() []PushMessageParamDO {
    return r._pushList
}
