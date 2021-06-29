package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推荐律师接口 API请求
alibaba.legal.suit.court.lawyer.push

为诉讼系统推荐律师
*/
type AlibabaLegalSuitCourtLawyerPushRequest struct {
    model.Params
    // 委托ID
    _entrustId   int64
    // 案件ID
    _suitId   int64
    // 推荐律师模型
    _lawyersModel   *LawyersModel
}

// 初始化AlibabaLegalSuitCourtLawyerPushRequest对象
func NewAlibabaLegalSuitCourtLawyerPushRequest() *AlibabaLegalSuitCourtLawyerPushRequest{
    return &AlibabaLegalSuitCourtLawyerPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtLawyerPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.lawyer.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtLawyerPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntrustId Setter
// 委托ID
func (r *AlibabaLegalSuitCourtLawyerPushRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalSuitCourtLawyerPushRequest) GetEntrustId() int64 {
    return r._entrustId
}
// SuitId Setter
// 案件ID
func (r *AlibabaLegalSuitCourtLawyerPushRequest) SetSuitId(_suitId int64) error {
    r._suitId = _suitId
    r.Set("suit_id", _suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalSuitCourtLawyerPushRequest) GetSuitId() int64 {
    return r._suitId
}
// LawyersModel Setter
// 推荐律师模型
func (r *AlibabaLegalSuitCourtLawyerPushRequest) SetLawyersModel(_lawyersModel *LawyersModel) error {
    r._lawyersModel = _lawyersModel
    r.Set("lawyers_model", _lawyersModel)
    return nil
}

// LawyersModel Getter
func (r AlibabaLegalSuitCourtLawyerPushRequest) GetLawyersModel() *LawyersModel {
    return r._lawyersModel
}
