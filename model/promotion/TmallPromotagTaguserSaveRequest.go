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
type TmallPromotagTaguserSaveAPIRequest struct {
    model.Params
    // 标签ID
    _tagId   int64
    // 买家昵称
    _nick   string
}

// 初始化TmallPromotagTaguserSaveAPIRequest对象
func NewTmallPromotagTaguserSaveRequest() *TmallPromotagTaguserSaveAPIRequest{
    return &TmallPromotagTaguserSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserSaveAPIRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.save"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagId Setter
// 标签ID
func (r *TmallPromotagTaguserSaveAPIRequest) SetTagId(_tagId int64) error {
    r._tagId = _tagId
    r.Set("tag_id", _tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTaguserSaveAPIRequest) GetTagId() int64 {
    return r._tagId
}
// Nick Setter
// 买家昵称
func (r *TmallPromotagTaguserSaveAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TmallPromotagTaguserSaveAPIRequest) GetNick() string {
    return r._nick
}
