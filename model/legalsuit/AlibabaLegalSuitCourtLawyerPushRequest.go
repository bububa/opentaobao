package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推荐律师接口 APIRequest
alibaba.legal.suit.court.lawyer.push

为诉讼系统推荐律师
*/
type AlibabaLegalSuitCourtLawyerPushRequest struct {
    model.Params

    // 委托ID
    entrustId   int64 

    // 案件ID
    suitId   int64 

    // 推荐律师模型
    lawyersModel   *LawyersModel 

}

func NewAlibabaLegalSuitCourtLawyerPushRequest() *AlibabaLegalSuitCourtLawyerPushRequest{
    return &AlibabaLegalSuitCourtLawyerPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitCourtLawyerPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.lawyer.push"
}

func (r AlibabaLegalSuitCourtLawyerPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitCourtLawyerPushRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalSuitCourtLawyerPushRequest) GetEntrustId() int64 {
    return r.entrustId
}

func (r *AlibabaLegalSuitCourtLawyerPushRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

func (r AlibabaLegalSuitCourtLawyerPushRequest) GetSuitId() int64 {
    return r.suitId
}

func (r *AlibabaLegalSuitCourtLawyerPushRequest) SetLawyersModel(lawyersModel *LawyersModel) error {
    r.lawyersModel = lawyersModel
    r.Set("lawyers_model", lawyersModel)
    return nil
}

func (r AlibabaLegalSuitCourtLawyerPushRequest) GetLawyersModel() *LawyersModel {
    return r.lawyersModel
}

