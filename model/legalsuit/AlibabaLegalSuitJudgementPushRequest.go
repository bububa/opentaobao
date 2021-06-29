package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送裁判登记信息给集团法务系统 API请求
alibaba.legal.suit.judgement.push

isv推送裁判登记信息给集团法务系统
*/
type AlibabaLegalSuitJudgementPushRequest struct {
    model.Params
    // 裁判登记信息
    _refereeRegistrationModel   *RefereeRegistrationModel
}

// 初始化AlibabaLegalSuitJudgementPushRequest对象
func NewAlibabaLegalSuitJudgementPushRequest() *AlibabaLegalSuitJudgementPushRequest{
    return &AlibabaLegalSuitJudgementPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitJudgementPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.judgement.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitJudgementPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefereeRegistrationModel Setter
// 裁判登记信息
func (r *AlibabaLegalSuitJudgementPushRequest) SetRefereeRegistrationModel(_refereeRegistrationModel *RefereeRegistrationModel) error {
    r._refereeRegistrationModel = _refereeRegistrationModel
    r.Set("referee_registration_model", _refereeRegistrationModel)
    return nil
}

// RefereeRegistrationModel Getter
func (r AlibabaLegalSuitJudgementPushRequest) GetRefereeRegistrationModel() *RefereeRegistrationModel {
    return r._refereeRegistrationModel
}
