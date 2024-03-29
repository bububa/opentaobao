package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTaguserJudgeAPIRequest 用户标签判断接口 API请求
// tmall.promotag.taguser.judge
//
// 查询用户是否有标签
type TmallPromotagTaguserJudgeAPIRequest struct {
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

// NewTmallPromotagTaguserJudgeRequest 初始化TmallPromotagTaguserJudgeAPIRequest对象
func NewTmallPromotagTaguserJudgeRequest() *TmallPromotagTaguserJudgeAPIRequest {
	return &TmallPromotagTaguserJudgeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallPromotagTaguserJudgeAPIRequest) Reset() {
	r._nick = ""
	r._ouid = ""
	r._openid = ""
	r._tagId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTaguserJudgeAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.taguser.judge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTaguserJudgeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPromotagTaguserJudgeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallPromotagTaguserJudgeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallPromotagTaguserJudgeRequest()
	},
}

// GetTmallPromotagTaguserJudgeRequest 从 sync.Pool 获取 TmallPromotagTaguserJudgeAPIRequest
func GetTmallPromotagTaguserJudgeAPIRequest() *TmallPromotagTaguserJudgeAPIRequest {
	return poolTmallPromotagTaguserJudgeAPIRequest.Get().(*TmallPromotagTaguserJudgeAPIRequest)
}

// ReleaseTmallPromotagTaguserJudgeAPIRequest 将 TmallPromotagTaguserJudgeAPIRequest 放入 sync.Pool
func ReleaseTmallPromotagTaguserJudgeAPIRequest(v *TmallPromotagTaguserJudgeAPIRequest) {
	v.Reset()
	poolTmallPromotagTaguserJudgeAPIRequest.Put(v)
}
