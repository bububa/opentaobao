package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTaguserJudgeAPIRequest 用户标签判断接口 API请求
// tmall.promotag.taguser.judge
//
// 查询用户是否有标签
type TmallPromotagTaguserJudgeAPIRequest struct {
	model.Params
	// 标签ID
	_tagId int64
	// 买家昵称
	_nick string
}

// NewTmallPromotagTaguserJudgeRequest 初始化TmallPromotagTaguserJudgeAPIRequest对象
func NewTmallPromotagTaguserJudgeRequest() *TmallPromotagTaguserJudgeAPIRequest {
	return &TmallPromotagTaguserJudgeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserJudgeAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.taguser.judge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserJudgeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTagId is TagId Setter
// 标签ID
func (r *TmallPromotagTaguserJudgeAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallPromotagTaguserJudgeAPIRequest) GetTagId() int64 {
	return r._tagId
}

// SetNick is Nick Setter
// 买家昵称
func (r *TmallPromotagTaguserJudgeAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallPromotagTaguserJudgeAPIRequest) GetNick() string {
	return r._nick
}
