package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家预存余额检查 API请求
alibaba.aliqin.flow.wallet.check.balance

检查商家CRM预存余额是否足够进行活动
*/
type AlibabaAliqinFlowWalletCheckBalanceRequest struct {
    model.Params
    // 检查金额档位id
    gradeId   string
}

// 初始化AlibabaAliqinFlowWalletCheckBalanceRequest对象
func NewAlibabaAliqinFlowWalletCheckBalanceRequest() *AlibabaAliqinFlowWalletCheckBalanceRequest{
    return &AlibabaAliqinFlowWalletCheckBalanceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletCheckBalanceRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.check.balance"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletCheckBalanceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GradeId Setter
// 检查金额档位id
func (r *AlibabaAliqinFlowWalletCheckBalanceRequest) SetGradeId(gradeId string) error {
    r.gradeId = gradeId
    r.Set("grade_id", gradeId)
    return nil
}

// GradeId Getter
func (r AlibabaAliqinFlowWalletCheckBalanceRequest) GetGradeId() string {
    return r.gradeId
}
