package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 API请求
taobao.tanx.creative.add

创意预审新增接口
*/
type TaobaoTanxCreativeAddRequest struct {
    model.Params
    // DSP的memberId
    _memberId   int64
    // dsp用户身份认证的TOKEN
    _token   string
    // 当前时间戳，1970-01-01后的秒数
    _signTime   int64
    // 创意id
    _creativeId   string
    // 广告类目 如果有多个，以逗号分隔
    _adboardType   string
    // 敏感词类目，多个的话以逗号分隔
    _sensitiveType   string
    // 创意代码
    _adboardData   string
    // 目标地址
    _destinationUrl   string
    // 创意尺寸,中间以小写字母x分隔
    _adboardSize   string
    // 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,3vast-linear（视频贴片） 4 vast-companion(视频伴随),5 Html5(无线banner或app)
    _creativePackageFormat   int64
}

// 初始化TaobaoTanxCreativeAddRequest对象
func NewTaobaoTanxCreativeAddRequest() *TaobaoTanxCreativeAddRequest{
    return &TaobaoTanxCreativeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxCreativeAddRequest) GetApiMethodName() string {
    return "taobao.tanx.creative.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxCreativeAddRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxCreativeAddRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativeAddRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxCreativeAddRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativeAddRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxCreativeAddRequest) GetSignTime() int64 {
    return r._signTime
}
// CreativeId Setter
// 创意id
func (r *TaobaoTanxCreativeAddRequest) SetCreativeId(_creativeId string) error {
    r._creativeId = _creativeId
    r.Set("creative_id", _creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoTanxCreativeAddRequest) GetCreativeId() string {
    return r._creativeId
}
// AdboardType Setter
// 广告类目 如果有多个，以逗号分隔
func (r *TaobaoTanxCreativeAddRequest) SetAdboardType(_adboardType string) error {
    r._adboardType = _adboardType
    r.Set("adboard_type", _adboardType)
    return nil
}

// AdboardType Getter
func (r TaobaoTanxCreativeAddRequest) GetAdboardType() string {
    return r._adboardType
}
// SensitiveType Setter
// 敏感词类目，多个的话以逗号分隔
func (r *TaobaoTanxCreativeAddRequest) SetSensitiveType(_sensitiveType string) error {
    r._sensitiveType = _sensitiveType
    r.Set("sensitive_type", _sensitiveType)
    return nil
}

// SensitiveType Getter
func (r TaobaoTanxCreativeAddRequest) GetSensitiveType() string {
    return r._sensitiveType
}
// AdboardData Setter
// 创意代码
func (r *TaobaoTanxCreativeAddRequest) SetAdboardData(_adboardData string) error {
    r._adboardData = _adboardData
    r.Set("adboard_data", _adboardData)
    return nil
}

// AdboardData Getter
func (r TaobaoTanxCreativeAddRequest) GetAdboardData() string {
    return r._adboardData
}
// DestinationUrl Setter
// 目标地址
func (r *TaobaoTanxCreativeAddRequest) SetDestinationUrl(_destinationUrl string) error {
    r._destinationUrl = _destinationUrl
    r.Set("destination_url", _destinationUrl)
    return nil
}

// DestinationUrl Getter
func (r TaobaoTanxCreativeAddRequest) GetDestinationUrl() string {
    return r._destinationUrl
}
// AdboardSize Setter
// 创意尺寸,中间以小写字母x分隔
func (r *TaobaoTanxCreativeAddRequest) SetAdboardSize(_adboardSize string) error {
    r._adboardSize = _adboardSize
    r.Set("adboard_size", _adboardSize)
    return nil
}

// AdboardSize Getter
func (r TaobaoTanxCreativeAddRequest) GetAdboardSize() string {
    return r._adboardSize
}
// CreativePackageFormat Setter
// 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,3vast-linear（视频贴片） 4 vast-companion(视频伴随),5 Html5(无线banner或app)
func (r *TaobaoTanxCreativeAddRequest) SetCreativePackageFormat(_creativePackageFormat int64) error {
    r._creativePackageFormat = _creativePackageFormat
    r.Set("creative_package_format", _creativePackageFormat)
    return nil
}

// CreativePackageFormat Getter
func (r TaobaoTanxCreativeAddRequest) GetCreativePackageFormat() int64 {
    return r._creativePackageFormat
}
