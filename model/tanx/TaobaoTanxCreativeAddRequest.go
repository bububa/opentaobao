package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 APIRequest
taobao.tanx.creative.add

创意预审新增接口
*/
type TaobaoTanxCreativeAddRequest struct {
    model.Params

    // DSP的memberId
    memberId   int64 

    // dsp用户身份认证的TOKEN
    token   string 

    // 当前时间戳，1970-01-01后的秒数
    signTime   int64 

    // 创意id
    creativeId   string 

    // 广告类目 如果有多个，以逗号分隔
    adboardType   string 

    // 敏感词类目，多个的话以逗号分隔
    sensitiveType   string 

    // 创意代码
    adboardData   string 

    // 目标地址
    destinationUrl   string 

    // 创意尺寸,中间以小写字母x分隔
    adboardSize   string 

    // 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,3vast-linear（视频贴片） 4 vast-companion(视频伴随),5 Html5(无线banner或app)
    creativePackageFormat   int64 

}

func NewTaobaoTanxCreativeAddRequest() *TaobaoTanxCreativeAddRequest{
    return &TaobaoTanxCreativeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxCreativeAddRequest) GetApiMethodName() string {
    return "taobao.tanx.creative.add"
}

func (r TaobaoTanxCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxCreativeAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxCreativeAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxCreativeAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxCreativeAddRequest) SetCreativeId(creativeId string) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetCreativeId() string {
    return r.creativeId
}

func (r *TaobaoTanxCreativeAddRequest) SetAdboardType(adboardType string) error {
    r.adboardType = adboardType
    r.Set("adboard_type", adboardType)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetAdboardType() string {
    return r.adboardType
}

func (r *TaobaoTanxCreativeAddRequest) SetSensitiveType(sensitiveType string) error {
    r.sensitiveType = sensitiveType
    r.Set("sensitive_type", sensitiveType)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetSensitiveType() string {
    return r.sensitiveType
}

func (r *TaobaoTanxCreativeAddRequest) SetAdboardData(adboardData string) error {
    r.adboardData = adboardData
    r.Set("adboard_data", adboardData)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetAdboardData() string {
    return r.adboardData
}

func (r *TaobaoTanxCreativeAddRequest) SetDestinationUrl(destinationUrl string) error {
    r.destinationUrl = destinationUrl
    r.Set("destination_url", destinationUrl)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetDestinationUrl() string {
    return r.destinationUrl
}

func (r *TaobaoTanxCreativeAddRequest) SetAdboardSize(adboardSize string) error {
    r.adboardSize = adboardSize
    r.Set("adboard_size", adboardSize)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetAdboardSize() string {
    return r.adboardSize
}

func (r *TaobaoTanxCreativeAddRequest) SetCreativePackageFormat(creativePackageFormat int64) error {
    r.creativePackageFormat = creativePackageFormat
    r.Set("creative_package_format", creativePackageFormat)
    return nil
}

func (r TaobaoTanxCreativeAddRequest) GetCreativePackageFormat() int64 {
    return r.creativePackageFormat
}

