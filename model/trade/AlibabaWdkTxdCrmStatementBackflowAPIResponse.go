package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达商家会员账单回流 API返回值 
alibaba.wdk.txd.crm.statement.backflow

淘鲜达商家会员账单回流
*/
type AlibabaWdkTxdCrmStatementBackflowAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTxdCrmStatementBackflowAPIResponseModel
}

// 淘鲜达商家会员账单回流 成功返回结果
type AlibabaWdkTxdCrmStatementBackflowAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_txd_crm_statement_backflow_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaWdkTxdCrmStatementBackflowApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
