package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝疫苗绑定接种人 APIRequest
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

func NewAlibabaHealthVaccinNoticeUserBindRequest() *AlibabaHealthVaccinNoticeUserBindRequest{
    return &AlibabaHealthVaccinNoticeUserBindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.user.bind"
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetBindUsers(bindUsers []AlipayVaccineUserBindDto) error {
    r.bindUsers = bindUsers
    r.Set("bind_users", bindUsers)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetBindUsers() []AlipayVaccineUserBindDto {
    return r.bindUsers
}

func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetOuterUserId(outerUserId string) error {
    r.outerUserId = outerUserId
    r.Set("outer_user_id", outerUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetOuterUserId() string {
    return r.outerUserId
}

func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetMobile() string {
    return r.mobile
}

func (r *AlibabaHealthVaccinNoticeUserBindRequest) SetAppChannel(appChannel string) error {
    r.appChannel = appChannel
    r.Set("app_channel", appChannel)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserBindRequest) GetAppChannel() string {
    return r.appChannel
}

