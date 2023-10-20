package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTaguserSaveAPIRequest 给用户打上优惠标签 API请求
// tmall.promotag.taguser.save
//
// 给用户载体打标
type TmallPromotagTaguserSaveAPIRequest struct {
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

// NewTmallPromotagTaguserSaveRequest 初始化TmallPromotagTaguserSaveAPIRequest对象
func NewTmallPromotagTaguserSaveRequest() *TmallPromotagTaguserSaveAPIRequest {
	return &TmallPromotagTaguserSaveAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallPromotagTaguserSaveAPIRequest) Reset() {
	r._nick = ""
	r._ouid = ""
	r._openid = ""
	r._tagId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserSaveAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.taguser.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPromotagTaguserSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TmallPromotagTaguserSaveAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallPromotagTaguserSaveAPIRequest) GetNick() string {
	return r._nick
}

// SetOuid is Ouid Setter
// 用户ID
func (r *TmallPromotagTaguserSaveAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallPromotagTaguserSaveAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// 用户ID
func (r *TmallPromotagTaguserSaveAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TmallPromotagTaguserSaveAPIRequest) GetOpenid() string {
	return r._openid
}

// SetTagId is TagId Setter
// 标签ID
func (r *TmallPromotagTaguserSaveAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallPromotagTaguserSaveAPIRequest) GetTagId() int64 {
	return r._tagId
}

var poolTmallPromotagTaguserSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallPromotagTaguserSaveRequest()
	},
}

// GetTmallPromotagTaguserSaveRequest 从 sync.Pool 获取 TmallPromotagTaguserSaveAPIRequest
func GetTmallPromotagTaguserSaveAPIRequest() *TmallPromotagTaguserSaveAPIRequest {
	return poolTmallPromotagTaguserSaveAPIRequest.Get().(*TmallPromotagTaguserSaveAPIRequest)
}

// ReleaseTmallPromotagTaguserSaveAPIRequest 将 TmallPromotagTaguserSaveAPIRequest 放入 sync.Pool
func ReleaseTmallPromotagTaguserSaveAPIRequest(v *TmallPromotagTaguserSaveAPIRequest) {
	v.Reset()
	poolTmallPromotagTaguserSaveAPIRequest.Put(v)
}
