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
type AlibabaHealthVaccinNoticeUserCreateAPIRequest struct {
    model.Params
    // 支付宝用户id
    _aliPayUserId   string
    // 外部渠道用户id
    _outerUserId   string
    // 用户电话号码
    _mobile   string
}

// 初始化AlibabaHealthVaccinNoticeUserCreateAPIRequest对象
func NewAlibabaHealthVaccinNoticeUserCreateRequest() *AlibabaHealthVaccinNoticeUserCreateAPIRequest{
    return &AlibabaHealthVaccinNoticeUserCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.user.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliPayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) SetAliPayUserId(_aliPayUserId string) error {
    r._aliPayUserId = _aliPayUserId
    r.Set("ali_pay_user_id", _aliPayUserId)
    return nil
}

// AliPayUserId Getter
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetAliPayUserId() string {
    return r._aliPayUserId
}
// OuterUserId Setter
// 外部渠道用户id
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) SetOuterUserId(_outerUserId string) error {
    r._outerUserId = _outerUserId
    r.Set("outer_user_id", _outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetOuterUserId() string {
    return r._outerUserId
}
// Mobile Setter
// 用户电话号码
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetMobile() string {
    return r._mobile
}
