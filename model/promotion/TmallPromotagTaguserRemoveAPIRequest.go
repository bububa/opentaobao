package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTaguserRemoveAPIRequest 给用户移除优惠标签 API请求
// tmall.promotag.taguser.remove
//
// 给用户载体去标
type TmallPromotagTaguserRemoveAPIRequest struct {
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

// NewTmallPromotagTaguserRemoveRequest 初始化TmallPromotagTaguserRemoveAPIRequest对象
func NewTmallPromotagTaguserRemoveRequest() *TmallPromotagTaguserRemoveAPIRequest {
	return &TmallPromotagTaguserRemoveAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallPromotagTaguserRemoveAPIRequest) Reset() {
	r._nick = ""
	r._ouid = ""
	r._openid = ""
	r._tagId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserRemoveAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.taguser.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPromotagTaguserRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TmallPromotagTaguserRemoveAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallPromotagTaguserRemoveAPIRequest) GetNick() string {
	return r._nick
}

// SetOuid is Ouid Setter
// 用户ID
func (r *TmallPromotagTaguserRemoveAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallPromotagTaguserRemoveAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// 用户ID
func (r *TmallPromotagTaguserRemoveAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TmallPromotagTaguserRemoveAPIRequest) GetOpenid() string {
	return r._openid
}

// SetTagId is TagId Setter
// 标签ID
func (r *TmallPromotagTaguserRemoveAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallPromotagTaguserRemoveAPIRequest) GetTagId() int64 {
	return r._tagId
}

var poolTmallPromotagTaguserRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallPromotagTaguserRemoveRequest()
	},
}

// GetTmallPromotagTaguserRemoveRequest 从 sync.Pool 获取 TmallPromotagTaguserRemoveAPIRequest
func GetTmallPromotagTaguserRemoveAPIRequest() *TmallPromotagTaguserRemoveAPIRequest {
	return poolTmallPromotagTaguserRemoveAPIRequest.Get().(*TmallPromotagTaguserRemoveAPIRequest)
}

// ReleaseTmallPromotagTaguserRemoveAPIRequest 将 TmallPromotagTaguserRemoveAPIRequest 放入 sync.Pool
func ReleaseTmallPromotagTaguserRemoveAPIRequest(v *TmallPromotagTaguserRemoveAPIRequest) {
	v.Reset()
	poolTmallPromotagTaguserRemoveAPIRequest.Put(v)
}
