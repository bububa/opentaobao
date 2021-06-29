package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝医疗健康疫苗用户创建 API请求
alibaba.health.vaccin.notice.user.create

支付宝医疗健康疫苗用户创建
*/
type AlibabaHealthVaccinNoticeUserCreateRequest struct {
    model.Params
    // 支付宝用户id
    aliPayUserId   string
    // 外部渠道用户id
    outerUserId   string
    // 用户电话号码
    mobile   string
}

// 初始化AlibabaHealthVaccinNoticeUserCreateRequest对象
func NewAlibabaHealthVaccinNoticeUserCreateRequest() *AlibabaHealthVaccinNoticeUserCreateRequest{
    return &AlibabaHealthVaccinNoticeUserCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.user.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliPayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeUserCreateRequest) SetAliPayUserId(aliPayUserId string) error {
    r.aliPayUserId = aliPayUserId
    r.Set("ali_pay_user_id", aliPayUserId)
    return nil
}

// AliPayUserId Getter
func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetAliPayUserId() string {
    return r.aliPayUserId
}
// OuterUserId Setter
// 外部渠道用户id
func (r *AlibabaHealthVaccinNoticeUserCreateRequest) SetOuterUserId(outerUserId string) error {
    r.outerUserId = outerUserId
    r.Set("outer_user_id", outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetOuterUserId() string {
    return r.outerUserId
}
// Mobile Setter
// 用户电话号码
func (r *AlibabaHealthVaccinNoticeUserCreateRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetMobile() string {
    return r.mobile
}
