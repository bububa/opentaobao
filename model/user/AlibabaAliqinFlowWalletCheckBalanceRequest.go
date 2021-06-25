package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家预存余额检查 APIRequest
alibaba.aliqin.flow.wallet.check.balance

检查商家CRM预存余额是否足够进行活动
*/
type AlibabaAliqinFlowWalletCheckBalanceRequest struct {
    model.Params

    // 检查金额档位id
    gradeId   string 

}

func NewAlibabaAliqinFlowWalletCheckBalanceRequest() *AlibabaAliqinFlowWalletCheckBalanceRequest{
    return &AlibabaAliqinFlowWalletCheckBalanceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletCheckBalanceRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.check.balance"
}

func (r AlibabaAliqinFlowWalletCheckBalanceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletCheckBalanceRequest) SetGradeId(gradeId string) error {
    r.gradeId = gradeId
    r.Set("grade_id", gradeId)
    return nil
}

func (r AlibabaAliqinFlowWalletCheckBalanceRequest) GetGradeId() string {
    return r.gradeId
}

