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
    alipayUserId   string
    // 绑定人信息list
    bindUsers   []AlipayVaccineUserBindDto
    // ISV 侧用户 ID
    outerUserId   string
    // 联系电话
    mobile   string
    // 用户来源：alipay或yl
    appChannel   string
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
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// BindUsers Setter
// 绑定人信息list
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetBindUsers(bindUsers []AlipayVaccineUserBindDto) error {
    r.bindUsers = bindUsers
    r.Set("bind_users", bindUsers)
    return nil
}

// BindUsers Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetBindUsers() []AlipayVaccineUserBindDto {
    return r.bindUsers
}
// OuterUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetOuterUserId(outerUserId string) error {
    r.outerUserId = outerUserId
    r.Set("outer_user_id", outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetOuterUserId() string {
    return r.outerUserId
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetMobile() string {
    return r.mobile
}
// AppChannel Setter
// 用户来源：alipay或yl
func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetAppChannel(appChannel string) error {
    r.appChannel = appChannel
    r.Set("app_channel", appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaHealthVaccinNoticeUserBindRequest) GetAppChannel() string {
    return r.appChannel
}
