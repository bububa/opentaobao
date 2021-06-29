package txcs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商账单确认 API返回值 
tmall.txcs.finance.bill.confirm

提供天猫超市外部合作商家：财务账单对账
*/
type TmallTxcsFinanceBillConfirmAPIResponse struct {
    model.CommonResponse
    TmallTxcsFinanceBillConfirmResponse
}

// 供应商账单确认 成功返回结果
type TmallTxcsFinanceBillConfirmResponse struct {
    XMLName xml.Name `xml:"tmall_txcs_finance_bill_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
