package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达商家会员账单回流 APIRequest
alibaba.wdk.txd.crm.statement.backflow

淘鲜达商家会员账单回流
*/
type AlibabaWdkTxdCrmStatementBackflowRequest struct {
    model.Params

    // 参数
    paramStatementBO   *StatementBo 

}

func NewAlibabaWdkTxdCrmStatementBackflowRequest() *AlibabaWdkTxdCrmStatementBackflowRequest{
    return &AlibabaWdkTxdCrmStatementBackflowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTxdCrmStatementBackflowRequest) GetApiMethodName() string {
    return "alibaba.wdk.txd.crm.statement.backflow"
}

func (r AlibabaWdkTxdCrmStatementBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTxdCrmStatementBackflowRequest) SetParamStatementBO(paramStatementBO *StatementBo) error {
    r.paramStatementBO = paramStatementBO
    r.Set("param_statement_b_o", paramStatementBO)
    return nil
}

func (r AlibabaWdkTxdCrmStatementBackflowRequest) GetParamStatementBO() *StatementBo {
    return r.paramStatementBO
}

