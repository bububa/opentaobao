package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户打上优惠标签 APIRequest
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

func NewTmallPromotagTaguserSaveRequest() *TmallPromotagTaguserSaveRequest{
    return &TmallPromotagTaguserSaveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotagTaguserSaveRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.save"
}

func (r TmallPromotagTaguserSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotagTaguserSaveRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

func (r TmallPromotagTaguserSaveRequest) GetTagId() int64 {
    return r.tagId
}

func (r *TmallPromotagTaguserSaveRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TmallPromotagTaguserSaveRequest) GetNick() string {
    return r.nick
}

