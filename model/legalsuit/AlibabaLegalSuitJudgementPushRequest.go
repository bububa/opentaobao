package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送裁判登记信息给集团法务系统 APIRequest
alibaba.legal.suit.judgement.push

isv推送裁判登记信息给集团法务系统
*/
type AlibabaLegalSuitJudgementPushRequest struct {
    model.Params

    // 裁判登记信息
    refereeRegistrationModel   *RefereeRegistrationModel 

}

func NewAlibabaLegalSuitJudgementPushRequest() *AlibabaLegalSuitJudgementPushRequest{
    return &AlibabaLegalSuitJudgementPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitJudgementPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.judgement.push"
}

func (r AlibabaLegalSuitJudgementPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitJudgementPushRequest) SetRefereeRegistrationModel(refereeRegistrationModel *RefereeRegistrationModel) error {
    r.refereeRegistrationModel = refereeRegistrationModel
    r.Set("referee_registration_model", refereeRegistrationModel)
    return nil
}

func (r AlibabaLegalSuitJudgementPushRequest) GetRefereeRegistrationModel() *RefereeRegistrationModel {
    return r.refereeRegistrationModel
}

