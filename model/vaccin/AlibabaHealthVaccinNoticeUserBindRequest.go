package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝疫苗绑定接种人 API请求
alibaba.health.vaccin.notice.user.bind

支付宝疫苗绑定接种人
*/
type AlibabaHealthVaccinNoticeUserBindRequest struct {
    model.Params
    // 支付宝ID
    _alipayUserId   string
    // 绑定人信息list
    _bindUsers   []AlipayVaccineUserBindDto
    // ISV 侧用户 ID
    _outerUserId   string
    // 联系电话
    _mobile   string
    // 用户来源：alipay或yl
    _appChannel   string
}

// 初始化AlibabaHealthVaccinNoticeUserBindRequest对象
func NewAlibabaHealthVaccinNoticeUserBindRequest() *AlibabaHealthVaccinNoticeUserBindRequest{
    return &AlibabaHealthVaccinNoticeUserBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.user.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝ID
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// BindUsers Setter
// 绑定人信息list
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetBindUsers(_bindUsers []AlipayVaccineUserBindDto) error {
    r._bindUsers = _bindUsers
    r.Set("bind_users", _bindUsers)
    return nil
}

// BindUsers Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetBindUsers() []AlipayVaccineUserBindDto {
    return r._bindUsers
}
// OuterUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetOuterUserId(_outerUserId string) error {
    r._outerUserId = _outerUserId
    r.Set("outer_user_id", _outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetOuterUserId() string {
    return r._outerUserId
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetMobile() string {
    return r._mobile
}
// AppChannel Setter
// 用户来源：alipay或yl
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetAppChannel(_appChannel string) error {
    r._appChannel = _appChannel
    r.Set("app_channel", _appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetAppChannel() string {
    return r._appChannel
}
