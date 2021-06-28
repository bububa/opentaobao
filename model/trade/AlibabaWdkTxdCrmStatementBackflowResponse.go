package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达商家会员账单回流 APIResponse
alibaba.wdk.txd.crm.statement.backflow

淘鲜达商家会员账单回流
*/
type AlibabaWdkTxdCrmStatementBackflowAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_txd_crm_statement_backflow_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkTxdCrmStatementBackflowApiResult `json:"result,omitempty" xml:"