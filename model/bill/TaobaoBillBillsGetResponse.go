package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账单明细数据(自研发商家专用) APIResponse
taobao.bill.bills.get

查询账单明细数据
*/
type TaobaoBillBillsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bill_bills_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 账单列表
    
    Bills   []Bill `json:"bills,omitempty" xml:"