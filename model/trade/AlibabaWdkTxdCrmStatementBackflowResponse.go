package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达商家会员账单回流 APIResponse
alibaba.wdk.txd.crm.statement.backflow

淘鲜达商家会员账单回流
*/
type AlibabaWdkTxdCrmStatementBackflowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkTxdCrmStatementBackflowResponse `json:"alibaba_wdk_txd_crm_statement_backflow_response,omitempty"`
}

type AlibabaWdkTxdCrmStatementBackflowResponse struct {

    // 结果
    Result   *AlibabaWdkTxdCrmStatementBackflowApiResult `json:"result,omitempty"`

}
