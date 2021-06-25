package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户移除优惠标签 APIRequest
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

func NewTmallPromotagTaguserRemoveRequest() *TmallPromotagTaguserRemoveRequest{
    return &TmallPromotagTaguserRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotagTaguserRemoveRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.remove"
}

func (r TmallPromotagTaguserRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotagTaguserRemoveRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

func (r TmallPromotagTaguserRemoveRequest) GetTagId() int64 {
    return r.tagId
}

func (r *TmallPromotagTaguserRemoveRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TmallPromotagTaguserRemoveRequest) GetNick() string {
    return r.nick
}

