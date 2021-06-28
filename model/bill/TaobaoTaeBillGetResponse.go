package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔账单明细 APIResponse
taobao.tae.bill.get

查询单笔账单明细
*/
type TaobaoTaeBillGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tae_bill_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 账单明细
    
    Bill   *BillDto `json:"bill,omitempty" xml:"