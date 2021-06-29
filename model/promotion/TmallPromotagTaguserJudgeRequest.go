package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户标签判断接口 API请求
tmall.promotag.taguser.judge

查询用户是否有标签
*/
type TmallPromotagTaguserJudgeRequest struct {
    model.Params
    // 标签ID
    tagId   int64
    // 买家昵称
    nick   string
}

// 初始化TmallPromotagTaguserJudgeRequest对象
func NewTmallPromotagTaguserJudgeRequest() *TmallPromotagTaguserJudgeRequest{
    return &TmallPromotagTaguserJudgeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserJudgeRequest) GetApiMethodName() string {
    return "tmall.promotag.taguser.judge"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserJudgeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagId Setter
// 标签ID
func (r *TmallPromotagTaguserJudgeRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTaguserJudgeRequest) GetTagId() int64 {
    return r.tagId
}
// Nick Setter
// 买家昵称
func (r *TmallPromotagTaguserJudgeRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TmallPromotagTaguserJudgeRequest) GetNick() string {
    return r.nick
}
