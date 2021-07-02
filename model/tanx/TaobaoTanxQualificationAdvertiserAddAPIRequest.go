package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxQualificationAdvertiserAddAPIRequest 新增广告主接口 API请求
// taobao.tanx.qualification.advertiser.add
//
// 外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主
type TaobaoTanxQualificationAdvertiserAddAPIRequest struct {
	model.Params
	// 广告主对象
	_advertisers []AdvertiserDto
	// dsp用户memberId
	_memberId int64
	// dsp用户验证token
	_token string
	// 从1970年到当前时间的秒
	_signTime int64
}

// NewTaobaoTanxQualificationAdvertiserAddRequest 初始化TaobaoTanxQualificationAdvertiserAddAPIRequest对象
func NewTaobaoTanxQualificationAdvertiserAddRequest() *TaobaoTanxQualificationAdvertiserAddAPIRequest {
	return &TaobaoTanxQualificationAdvertiserAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationAdvertiserAddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.advertiser.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationAdvertiserAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAdvertisers is Advertisers Setter
// 广告主对象
func (r *TaobaoTanxQualificationAdvertiserAddAPIRequest) SetAdvertisers(_advertisers []AdvertiserDto) error {
	r._advertisers = _advertisers
	r.Set("advertisers", _advertisers)
	return nil
}

// GetAdvertisers Advertisers Getter
func (r TaobaoTanxQualificationAdvertiserAddAPIRequest) GetAdvertisers() []AdvertiserDto {
	return r._advertisers
}

// SetMemberId is MemberId Setter
// dsp用户memberId
func (r *TaobaoTanxQualificationAdvertiserAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxQualificationAdvertiserAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetToken is Token Setter
// dsp用户验证token
func (r *TaobaoTanxQualificationAdvertiserAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxQualificationAdvertiserAddAPIRequest) GetToken() string {
	return r._token
}

// SetSignTime is SignTime Setter
// 从1970年到当前时间的秒
func (r *TaobaoTanxQualificationAdvertiserAddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxQualificationAdvertiserAddAPIRequest) GetSignTime() int64 {
	return r._signTime
}
