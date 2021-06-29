package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或者保存管辖信息 APIRequest
alibaba.legal.suit.domination.push

ISV推送管辖信息到诉讼平台
*/
type AlibabaLegalSuitDominationPushRequest struct {
    model.Params

    // 管辖信息
    dominationDissentModel   *DominationDissentModel 

}

func NewAlibabaLegalSuitDominationPushRequest() *AlibabaLegalSuitDominationPushRequest{
    return &AlibabaLegalSuitDominationPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitDominationPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.domination.push"
}

func (r AlibabaLegalSuitDominationPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitDominationPushRequest) SetDominationDissentModel(dominationDissentModel *DominationDissentModel) error {
    r.dominationDissentModel = dominationDissentModel
    r.Set("domination_dissent_model", dominationDissentModel)
    return nil
}

func (r AlibabaLegalSuitDominationPushRequest) GetDominationDissentModel() *DominationDissentModel {
    return r.dominationDissentModel
}

