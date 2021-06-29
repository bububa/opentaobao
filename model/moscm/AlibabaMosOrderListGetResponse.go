package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询订单交易 APIResponse
alibaba.mos.order.list.get

批量查询交易信息
*/
type AlibabaMosOrderListGetAPIResponse struct {
    model.CommonResponse
    AlibabaMosOrderListGetResponse
}

type AlibabaMosOrderListGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_order_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *AlibabaMosOrderListGetResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
