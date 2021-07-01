package txcs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫超市外部商家财务账单信息查询 API返回值 
tmall.txcs.finance.bill.query

提供天猫超市外部合作商家财务账单对账
*/
type TmallTxcsFinanceBillQueryAPIResponse struct {
    model.CommonResponse
    TmallTxcsFinanceBillQueryAPIResponseModel
}

// 天猫超市外部商家财务账单信息查询 成功返回结果
type TmallTxcsFinanceBillQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_txcs_finance_bill_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
