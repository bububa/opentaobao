package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberJoinurlGetAPIRequest 会员入会地址获取 API请求
// taobao.crm.member.joinurl.get
//
// 会员入会地址获取
type TaobaoCrmMemberJoinurlGetAPIRequest struct {
	model.Params
	// 回调url
	_callbackUrl string
	// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 isvapp可录入的其他参数，活动ID：activityId三方招募来源：entrance
	_extraInfo string
}

// NewTaobaoCrmMemberJoinurlGetRequest 初始化TaobaoCrmMemberJoinurlGetAPIRequest对象
func NewTaobaoCrmMemberJoinurlGetRequest() *TaobaoCrmMemberJoinurlGetAPIRequest {
	return &TaobaoCrmMemberJoinurlGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMemberJoinurlGetAPIRequest) Reset() {
	r._callbackUrl = ""
	r._extraInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberJoinurlGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.member.joinurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberJoinurlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMemberJoinurlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackUrl is CallbackUrl Setter
// 回调url
func (r *TaobaoCrmMemberJoinurlGetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r TaobaoCrmMemberJoinurlGetAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetExtraInfo is ExtraInfo Setter
// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 isvapp可录入的其他参数，活动ID：activityId三方招募来源：entrance
func (r *TaobaoCrmMemberJoinurlGetAPIRequest) SetExtraInfo(_extraInfo string) error {
	r._extraInfo = _extraInfo
	r.Set("extra_info", _extraInfo)
	return nil
}

// GetExtraInfo ExtraInfo Getter
func (r TaobaoCrmMemberJoinurlGetAPIRequest) GetExtraInfo() string {
	return r._extraInfo
}

var poolTaobaoCrmMemberJoinurlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMemberJoinurlGetRequest()
	},
}

// GetTaobaoCrmMemberJoinurlGetRequest 从 sync.Pool 获取 TaobaoCrmMemberJoinurlGetAPIRequest
func GetTaobaoCrmMemberJoinurlGetAPIRequest() *TaobaoCrmMemberJoinurlGetAPIRequest {
	return poolTaobaoCrmMemberJoinurlGetAPIRequest.Get().(*TaobaoCrmMemberJoinurlGetAPIRequest)
}

// ReleaseTaobaoCrmMemberJoinurlGetAPIRequest 将 TaobaoCrmMemberJoinurlGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMemberJoinurlGetAPIRequest(v *TaobaoCrmMemberJoinurlGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmMemberJoinurlGetAPIRequest.Put(v)
}
