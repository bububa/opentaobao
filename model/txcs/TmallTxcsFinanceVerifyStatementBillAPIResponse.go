package txcs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商核销单录入 API返回值 
tmall.txcs.finance.verify.statement.bill

供应商核销单录入
*/
type TmallTxcsFinanceVerifyStatementBillAPIResponse struct {
    model.CommonResponse
    TmallTxcsFinanceVerifyStatementBillAPIResponseModel
}

// 供应商核销单录入 成功返回结果
type TmallTxcsFinanceVerifyStatementBillAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_txcs_finance_verify_statement_bill_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
