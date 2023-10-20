package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallpromotagtaguserjudgeAPIRequest 用户标签判断接口 API请求
// tmall.promotag.taguser.judge
//
// 查询用户是否有标签
type TmallpromotagtaguserjudgeAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 买家ID
	_ouid string
	// 买家ID
	_openid string
	// 标签ID
	_tagId int64
}

// NewTmallpromotagtaguserjudgeRequest 初始化TmallpromotagtaguserjudgeAPIRequest对象
func NewTmallpromotagtaguserjudgeRequest() *TmallpromotagtaguserjudgeAPIRequest {
	return &TmallpromotagtaguserjudgeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallpromotagtaguserjudgeAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.taguser.judge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallpromotagtaguserjudgeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallpromotagtaguserjudgeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TmallpromotagtaguserjudgeAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallpromotagtaguserjudgeAPIRequest) GetNick() string {
	return r._nick
}

// SetOuid is Ouid Setter
// 买家ID
func (r *TmallpromotagtaguserjudgeAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallpromotagtaguserjudgeAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// 买家ID
func (r *TmallpromotagtaguserjudgeAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TmallpromotagtaguserjudgeAPIRequest) GetOpenid() string {
	return r._openid
}

// SetTagId is TagId Setter
// 标签ID
func (r *TmallpromotagtaguserjudgeAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallpromotagtaguserjudgeAPIRequest) GetTagId() int64 {
	return r._tagId
}
