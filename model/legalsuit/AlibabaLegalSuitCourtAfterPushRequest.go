package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或者新增庭后信息 APIRequest
alibaba.legal.suit.court.after.push

供外部ISV供应商 推送庭后信息给集团诉讼
*/
type AlibabaLegalSuitCourtAfterPushRequest struct {
    model.Params

    // 庭后信息
    afterCourtInfoModel   *AfterCourtInfoModel 

}

func NewAlibabaLegalSuitCourtAfterPushRequest() *AlibabaLegalSuitCourtAfterPushRequest{
    return &AlibabaLegalSuitCourtAfterPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitCourtAfterPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.after.push"
}

func (r AlibabaLegalSuitCourtAfterPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitCourtAfterPushRequest) SetAfterCourtInfoModel(afterCourtInfoModel *AfterCourtInfoModel) error {
    r.afterCourtInfoModel = afterCourtInfoModel
    r.Set("after_court_info_model", afterCourtInfoModel)
    return nil
}

func (r AlibabaLegalSuitCourtAfterPushRequest) GetAfterCourtInfoModel() *AfterCourtInfoModel {
    return r.afterCourtInfoModel
}

