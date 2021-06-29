package txcs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商核销单录入 APIResponse
tmall.txcs.finance.verify.statement.bill

供应商核销单录入
*/
type TmallTxcsFinanceVerifyStatementBillAPIResponse struct {
    model.CommonResponse
    TmallTxcsFinanceVerifyStatementBillResponse
}

type TmallTxcsFinanceVerifyStatementBillResponse struct {
    XMLName xml.Name `xml:"tmall_txcs_finance_verify_statement_bill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
