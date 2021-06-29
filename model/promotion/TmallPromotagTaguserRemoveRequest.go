package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户移除优惠标签 API请求
tmall.promotag.taguser.remove

给用户载体去标
*/
type TmallPromotagTaguserRemoveRequest struct {
    model.Params
    // 标签ID
    tagId   int64
    // 买家昵称
    nick   string
}

// 初始化TmallPromotagTaguserRemoveRequest对象
func NewTmallPromotagTaguserRemoveRequest() *TmallPromotagTaguserRemoveRequest{
    return &TmallPromotagTaguserRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserRemoveRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.remove"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagId Setter
// 标签ID
func (r *TmallPromotagTaguserRemoveRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTaguserRemoveRequest) GetTagId() int64 {
    return r.tagId
}
// Nick Setter
// 买家昵称
func (r *TmallPromotagTaguserRemoveRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TmallPromotagTaguserRemoveRequest) GetNick() string {
    return r.nick
}
