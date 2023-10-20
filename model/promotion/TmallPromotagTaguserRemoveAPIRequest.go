package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallpromotagtaguserremoveAPIRequest 给用户移除优惠标签 API请求
// tmall.promotag.taguser.remove
//
// 给用户载体去标
type TmallpromotagtaguserremoveAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 用户ID
	_ouid string
	// 用户ID
	_openid string
	// 标签ID
	_tagId int64
}

// NewTmallpromotagtaguserremoveRequest 初始化TmallpromotagtaguserremoveAPIRequest对象
func NewTmallpromotagtaguserremoveRequest() *TmallpromotagtaguserremoveAPIRequest {
	return &TmallpromotagtaguserremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallpromotagtaguserremoveAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.taguser.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallpromotagtaguserremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallpromotagtaguserremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TmallpromotagtaguserremoveAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallpromotagtaguserremoveAPIRequest) GetNick() string {
	return r._nick
}

// SetOuid is Ouid Setter
// 用户ID
func (r *TmallpromotagtaguserremoveAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallpromotagtaguserremoveAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// 用户ID
func (r *TmallpromotagtaguserremoveAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TmallpromotagtaguserremoveAPIRequest) GetOpenid() string {
	return r._openid
}

// SetTagId is TagId Setter
// 标签ID
func (r *TmallpromotagtaguserremoveAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallpromotagtaguserremoveAPIRequest) GetTagId() int64 {
	return r._tagId
}
