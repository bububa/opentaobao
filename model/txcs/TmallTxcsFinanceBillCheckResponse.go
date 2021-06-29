package txcs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫超市外部商家财务账单对账 APIResponse
tmall.txcs.finance.bill.check

提供天猫超市外部合作商家财务账单对账
*/
type TmallTxcsFinanceBillCheckAPIResponse struct {
    model.CommonResponse
    TmallTxcsFinanceBillCheckResponse
}

type TmallTxcsFinanceBillCheckResponse struct {
    XMLName xml.Name `xml:"tmall_txcs_finance_bill_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果
    
    Result   *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
