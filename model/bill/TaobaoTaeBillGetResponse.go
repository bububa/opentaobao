package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔账单明细 API返回值 
taobao.tae.bill.get

查询单笔账单明细
*/
type TaobaoTaeBillGetAPIResponse struct {
    model.CommonResponse
    TaobaoTaeBillGetResponse
}

// tae查询单笔账单明细 成功返回结果
type TaobaoTaeBillGetResponse struct {
    XMLName xml.Name `xml:"tae_bill_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 账单明细
    Bill   *BillDTO `json:"bill,omitempty" xml:"bill,omitempty"`
}
