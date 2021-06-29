package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询交易退货信息 APIResponse
alibaba.mos.order.refund.list.get

批量查询多个退货单的退货明细
*/
type AlibabaMosOrderRefundListGetAPIResponse struct {
    model.CommonResponse
    AlibabaMosOrderRefundListGetResponse
}

type AlibabaMosOrderRefundListGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_order_refund_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *AlibabaMosOrderRefundListGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
