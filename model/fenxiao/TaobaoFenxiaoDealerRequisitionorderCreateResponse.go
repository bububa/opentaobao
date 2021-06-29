package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建经销采购申请 APIResponse
taobao.fenxiao.dealer.requisitionorder.create

创建经销采购申请
*/
type TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoDealerRequisitionorderCreateResponse
}

type TaobaoFenxiaoDealerRequisitionorderCreateResponse struct {
    XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 经销采购申请编号
    
    DealerOrderId   int64 `json:"dealer_order_id,omitempty" xml:"dealer_order_id,omitempty"`

    
}
