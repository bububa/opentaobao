package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取裁判登记信息 APIRequest
alibaba.legal.suit.judgement.get

供ISV供应商获取集团法务系统的裁判登记信息
*/
type AlibabaLegalSuitJudgementGetRequest struct {
    model.Params

    // 案件id
    suitId   int64 

}

func NewAlibabaLegalSuitJudgementGetRequest() *AlibabaLegalSuitJudgementGetRequest{
    return &AlibabaLegalSuitJudgementGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitJudgementGetRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.judgement.get"
}

func (r AlibabaLegalSuitJudgementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitJudgementGetRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

func (r AlibabaLegalSuitJudgementGetRequest) GetSuitId() int64 {
    return r.suitId
}

