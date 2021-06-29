package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或保存庭前信息 APIRequest
alibaba.legal.suit.court.before.push

更新或者保存庭前信息
*/
type AlibabaLegalSuitCourtBeforePushRequest struct {
    model.Params

    // 庭前信息
    beforeCourtModel   *BeforeCourtModel 

}

func NewAlibabaLegalSuitCourtBeforePushRequest() *AlibabaLegalSuitCourtBeforePushRequest{
    return &AlibabaLegalSuitCourtBeforePushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitCourtBeforePushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.before.push"
}

func (r AlibabaLegalSuitCourtBeforePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitCourtBeforePushRequest) SetBeforeCourtModel(beforeCourtModel *BeforeCourtModel) error {
    r.beforeCourtModel = beforeCourtModel
    r.Set("before_court_model", beforeCourtModel)
    return nil
}

func (r AlibabaLegalSuitCourtBeforePushRequest) GetBeforeCourtModel() *BeforeCourtModel {
    return r.beforeCourtModel
}

