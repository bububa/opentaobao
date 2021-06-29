package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取裁判登记信息 API请求
alibaba.legal.suit.judgement.get

供ISV供应商获取集团法务系统的裁判登记信息
*/
type AlibabaLegalSuitJudgementGetRequest struct {
    model.Params
    // 案件id
    suitId   int64
}

// 初始化AlibabaLegalSuitJudgementGetRequest对象
func NewAlibabaLegalSuitJudgementGetRequest() *AlibabaLegalSuitJudgementGetRequest{
    return &AlibabaLegalSuitJudgementGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitJudgementGetRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.judgement.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitJudgementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SuitId Setter
// 案件id
func (r *AlibabaLegalSuitJudgementGetRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalSuitJudgementGetRequest) GetSuitId() int64 {
    return r.suitId
}
