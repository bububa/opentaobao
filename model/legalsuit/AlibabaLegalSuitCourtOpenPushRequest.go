package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开庭信息推送接口 APIRequest
alibaba.legal.suit.court.open.push

供ISV推送开庭信息给集团诉讼
*/
type AlibabaLegalSuitCourtOpenPushRequest struct {
    model.Params

    // 开庭信息
    courtInfoModel   *CourtInfoModel 

}

func NewAlibabaLegalSuitCourtOpenPushRequest() *AlibabaLegalSuitCourtOpenPushRequest{
    return &AlibabaLegalSuitCourtOpenPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitCourtOpenPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.open.push"
}

func (r AlibabaLegalSuitCourtOpenPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitCourtOpenPushRequest) SetCourtInfoModel(courtInfoModel *CourtInfoModel) error {
    r.courtInfoModel = courtInfoModel
    r.Set("court_info_model", courtInfoModel)
    return nil
}

func (r AlibabaLegalSuitCourtOpenPushRequest) GetCourtInfoModel() *CourtInfoModel {
    return r.courtInfoModel
}

