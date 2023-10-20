package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxcreativesgetAPIRequest 批量获取DSP用户的创意审核结果 API请求
// taobao.tanx.creatives.get
//
// 批量获取DSP用户的创意审核结果
type TaobaotanxcreativesgetAPIRequest struct {
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

// NewTaobaotanxcreativesgetRequest 初始化TaobaotanxcreativesgetAPIRequest对象
func NewTaobaotanxcreativesgetRequest() *TaobaotanxcreativesgetAPIRequest {
	return &TaobaotanxcreativesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxcreativesgetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.creatives.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxcreativesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxcreativesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaotanxcreativesgetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxcreativesgetAPIRequest) GetToken() string {
	return r._token
}

// SetStatus is Status Setter
// 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
func (r *TaobaotanxcreativesgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaotanxcreativesgetAPIRequest) GetStatus() string {
	return r._status
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaotanxcreativesgetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxcreativesgetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaotanxcreativesgetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxcreativesgetAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetPage is Page Setter
// 分页的页码(第一页为1)
func (r *TaobaotanxcreativesgetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaotanxcreativesgetAPIRequest) GetPage() int64 {
	return r._page
}

// SetType is Type Setter
// 所选创意的类型。1--&gt;普通类型, 2--&gt;视频贴片, 0 --&gt;优先查询普通类型,无结果则查询视频贴片类型
func (r *TaobaotanxcreativesgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaotanxcreativesgetAPIRequest) GetType() int64 {
	return r._type
}
