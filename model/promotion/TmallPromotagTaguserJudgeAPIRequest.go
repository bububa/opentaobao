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
	// 买家ID
	_ouid string
	// 买家ID
	_openid string
	// 昵称
	_nick string
	// 标签ID
	_tagId int64
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

// SetOuid is Ouid Setter
// 买家ID
func (r *TmallPromotagTaguserJudgeAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallPromotagTaguserJudgeAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// 买家ID
func (r *TmallPromotagTaguserJudgeAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TmallPromotagTaguserJudgeAPIRequest) GetOpenid() string {
	return r._openid
}

// SetNick is Nick Setter
// 昵称
func (r *TmallPromotagTaguserJudgeAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallPromotagTaguserJudgeAPIRequest) GetNick() string {
	return r._nick
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
