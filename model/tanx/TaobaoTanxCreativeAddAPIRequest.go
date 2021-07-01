package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxCreativeAddAPIRequest
创意预审新增接口 API请求
taobao.tanx.creative.add

创意预审新增接口 */
type TaobaoTanxCreativeAddAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 创意id
	_creativeId string
	// 广告类目 如果有多个，以逗号分隔
	_adboardType string
	// 敏感词类目，多个的话以逗号分隔
	_sensitiveType string
	// 创意代码
	_adboardData string
	// 目标地址
	_destinationUrl string
	// 创意尺寸,中间以小写字母x分隔
	_adboardSize string
	// 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,3vast-linear（视频贴片） 4 vast-companion(视频伴随),5 Html5(无线banner或app)
	_creativePackageFormat int64
}

// NewTaobaoTanxCreativeAddRequest 初始化TaobaoTanxCreativeAddAPIRequest对象
func NewTaobaoTanxCreativeAddRequest() *TaobaoTanxCreativeAddAPIRequest {
	return &TaobaoTanxCreativeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxCreativeAddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxCreativeAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MemberId Setter
// DSP的memberId
func (r *TaobaoTanxCreativeAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// Get MemberId Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// Set is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativeAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetToken() string {
	return r._token
}

// Set is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativeAddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// Get SignTime Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// Set is CreativeId Setter
// 创意id
func (r *TaobaoTanxCreativeAddAPIRequest) SetCreativeId(_creativeId string) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// Get CreativeId Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetCreativeId() string {
	return r._creativeId
}

// Set is AdboardType Setter
// 广告类目 如果有多个，以逗号分隔
func (r *TaobaoTanxCreativeAddAPIRequest) SetAdboardType(_adboardType string) error {
	r._adboardType = _adboardType
	r.Set("adboard_type", _adboardType)
	return nil
}

// Get AdboardType Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetAdboardType() string {
	return r._adboardType
}

// Set is SensitiveType Setter
// 敏感词类目，多个的话以逗号分隔
func (r *TaobaoTanxCreativeAddAPIRequest) SetSensitiveType(_sensitiveType string) error {
	r._sensitiveType = _sensitiveType
	r.Set("sensitive_type", _sensitiveType)
	return nil
}

// Get SensitiveType Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetSensitiveType() string {
	return r._sensitiveType
}

// Set is AdboardData Setter
// 创意代码
func (r *TaobaoTanxCreativeAddAPIRequest) SetAdboardData(_adboardData string) error {
	r._adboardData = _adboardData
	r.Set("adboard_data", _adboardData)
	return nil
}

// Get AdboardData Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetAdboardData() string {
	return r._adboardData
}

// Set is DestinationUrl Setter
// 目标地址
func (r *TaobaoTanxCreativeAddAPIRequest) SetDestinationUrl(_destinationUrl string) error {
	r._destinationUrl = _destinationUrl
	r.Set("destination_url", _destinationUrl)
	return nil
}

// Get DestinationUrl Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetDestinationUrl() string {
	return r._destinationUrl
}

// Set is AdboardSize Setter
// 创意尺寸,中间以小写字母x分隔
func (r *TaobaoTanxCreativeAddAPIRequest) SetAdboardSize(_adboardSize string) error {
	r._adboardSize = _adboardSize
	r.Set("adboard_size", _adboardSize)
	return nil
}

// Get AdboardSize Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetAdboardSize() string {
	return r._adboardSize
}

// Set is CreativePackageFormat Setter
// 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,3vast-linear（视频贴片） 4 vast-companion(视频伴随),5 Html5(无线banner或app)
func (r *TaobaoTanxCreativeAddAPIRequest) SetCreativePackageFormat(_creativePackageFormat int64) error {
	r._creativePackageFormat = _creativePackageFormat
	r.Set("creative_package_format", _creativePackageFormat)
	return nil
}

// Get CreativePackageFormat Getter
func (r TaobaoTanxCreativeAddAPIRequest) GetCreativePackageFormat() int64 {
	return r._creativePackageFormat
}
