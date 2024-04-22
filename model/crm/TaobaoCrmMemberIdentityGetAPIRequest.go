package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberIdentityGetAPIRequest 会员身份识别 API请求
// taobao.crm.member.identity.get
//
// 用来识别该用户是否是商家会员
type TaobaoCrmMemberIdentityGetAPIRequest struct {
	model.Params
	// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
	_extraInfo string
	// 混淆昵称，
	_mixNick string
	// 明文nick，可不填，直接填混淆昵称
	_nick string
	// open_id
	_openId string
}

// NewTaobaoCrmMemberIdentityGetRequest 初始化TaobaoCrmMemberIdentityGetAPIRequest对象
func NewTaobaoCrmMemberIdentityGetRequest() *TaobaoCrmMemberIdentityGetAPIRequest {
	return &TaobaoCrmMemberIdentityGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMemberIdentityGetAPIRequest) Reset() {
	r._extraInfo = ""
	r._mixNick = ""
	r._nick = ""
	r._openId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.member.identity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraInfo is ExtraInfo Setter
// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
func (r *TaobaoCrmMemberIdentityGetAPIRequest) SetExtraInfo(_extraInfo string) error {
	r._extraInfo = _extraInfo
	r.Set("extra_info", _extraInfo)
	return nil
}

// GetExtraInfo ExtraInfo Getter
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetExtraInfo() string {
	return r._extraInfo
}

// SetMixNick is MixNick Setter
// 混淆昵称，
func (r *TaobaoCrmMemberIdentityGetAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetMixNick() string {
	return r._mixNick
}

// SetNick is Nick Setter
// 明文nick，可不填，直接填混淆昵称
func (r *TaobaoCrmMemberIdentityGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetNick() string {
	return r._nick
}

// SetOpenId is OpenId Setter
// open_id
func (r *TaobaoCrmMemberIdentityGetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoCrmMemberIdentityGetAPIRequest) GetOpenId() string {
	return r._openId
}

var poolTaobaoCrmMemberIdentityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMemberIdentityGetRequest()
	},
}

// GetTaobaoCrmMemberIdentityGetRequest 从 sync.Pool 获取 TaobaoCrmMemberIdentityGetAPIRequest
func GetTaobaoCrmMemberIdentityGetAPIRequest() *TaobaoCrmMemberIdentityGetAPIRequest {
	return poolTaobaoCrmMemberIdentityGetAPIRequest.Get().(*TaobaoCrmMemberIdentityGetAPIRequest)
}

// ReleaseTaobaoCrmMemberIdentityGetAPIRequest 将 TaobaoCrmMemberIdentityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMemberIdentityGetAPIRequest(v *TaobaoCrmMemberIdentityGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmMemberIdentityGetAPIRequest.Put(v)
}
