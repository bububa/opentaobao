package tanx

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxCreativesGetAPIRequest 批量获取DSP用户的创意审核结果 API请求
// taobao.tanx.creatives.get
//
// 批量获取DSP用户的创意审核结果
type TaobaoTanxCreativesGetAPIRequest struct {
	model.Params
	// dsp用户身份认证的TOKEN
	_token string
	// 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
	_status string
	// DSP的memberId
	_memberId int64
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 分页的页码(第一页为1)
	_page int64
	// 所选创意的类型。1-->普通类型, 2-->视频贴片, 0 -->优先查询普通类型,无结果则查询视频贴片类型
	_type int64
}

// NewTaobaoTanxCreativesGetRequest 初始化TaobaoTanxCreativesGetAPIRequest对象
func NewTaobaoTanxCreativesGetRequest() *TaobaoTanxCreativesGetAPIRequest {
	return &TaobaoTanxCreativesGetAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTanxCreativesGetAPIRequest) Reset() {
	r._token = ""
	r._status = ""
	r._memberId = 0
	r._signTime = 0
	r._page = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxCreativesGetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.creatives.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxCreativesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTanxCreativesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativesGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxCreativesGetAPIRequest) GetToken() string {
	return r._token
}

// SetStatus is Status Setter
// 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
func (r *TaobaoTanxCreativesGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoTanxCreativesGetAPIRequest) GetStatus() string {
	return r._status
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaoTanxCreativesGetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxCreativesGetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativesGetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxCreativesGetAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetPage is Page Setter
// 分页的页码(第一页为1)
func (r *TaobaoTanxCreativesGetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaoTanxCreativesGetAPIRequest) GetPage() int64 {
	return r._page
}

// SetType is Type Setter
// 所选创意的类型。1--&gt;普通类型, 2--&gt;视频贴片, 0 --&gt;优先查询普通类型,无结果则查询视频贴片类型
func (r *TaobaoTanxCreativesGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoTanxCreativesGetAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoTanxCreativesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTanxCreativesGetRequest()
	},
}

// GetTaobaoTanxCreativesGetRequest 从 sync.Pool 获取 TaobaoTanxCreativesGetAPIRequest
func GetTaobaoTanxCreativesGetAPIRequest() *TaobaoTanxCreativesGetAPIRequest {
	return poolTaobaoTanxCreativesGetAPIRequest.Get().(*TaobaoTanxCreativesGetAPIRequest)
}

// ReleaseTaobaoTanxCreativesGetAPIRequest 将 TaobaoTanxCreativesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTanxCreativesGetAPIRequest(v *TaobaoTanxCreativesGetAPIRequest) {
	v.Reset()
	poolTaobaoTanxCreativesGetAPIRequest.Put(v)
}
