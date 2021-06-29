package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户打上优惠标签 API请求
tmall.promotag.taguser.save

给用户载体打标
*/
type TmallPromotagTaguserSaveRequest struct {
    model.Params
    // 标签ID
    tagId   int64
    // 买家昵称
    nick   string
}

// 初始化TmallPromotagTaguserSaveRequest对象
func NewTmallPromotagTaguserSaveRequest() *TmallPromotagTaguserSaveRequest{
    return &TmallPromotagTaguserSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserSaveRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.save"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagId Setter
// 标签ID
func (r *TmallPromotagTaguserSaveRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTaguserSaveRequest) GetTagId() int64 {
    return r.tagId
}
// Nick Setter
// 买家昵称
func (r *TmallPromotagTaguserSaveRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TmallPromotagTaguserSaveRequest) GetNick() string {
    return r.nick
}
