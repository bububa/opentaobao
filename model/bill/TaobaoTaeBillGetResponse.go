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
    TaobaoTaeBillGetResponse
}

type TaobaoTaeBillGetResponse struct {
    XMLName xml.Name `xml:"tae_bill_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 账单明细
    
    Bill   *BillDto `json:"bill,omitempty" xml:"bill,omitempty"`

    
}
