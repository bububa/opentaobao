package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝医疗健康疫苗用户创建 APIRequest
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

func NewAlibabaHealthVaccinNoticeUserCreateRequest() *AlibabaHealthVaccinNoticeUserCreateRequest{
    return &AlibabaHealthVaccinNoticeUserCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.user.create"
}

func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeUserCreateRequest) SetAliPayUserId(aliPayUserId string) error {
    r.aliPayUserId = aliPayUserId
    r.Set("ali_pay_user_id", aliPayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetAliPayUserId() string {
    return r.aliPayUserId
}

func (r *AlibabaHealthVaccinNoticeUserCreateRequest) SetOuterUserId(outerUserId string) error {
    r.outerUserId = outerUserId
    r.Set("outer_user_id", outerUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetOuterUserId() string {
    return r.outerUserId
}

func (r *AlibabaHealthVaccinNoticeUserCreateRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaHealthVaccinNoticeUserCreateRequest) GetMobile() string {
    return r.mobile
}

