package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过物流订单编号查询物流信息 APIResponse
taobao.wlb.tmsorder.query

通过物流订单编号分页查询物流信息
*/
type TaobaoWlbTmsorderQueryAPIResponse struct {
    model.CommonResponse
    TaobaoWlbTmsorderQueryResponse
}

type TaobaoWlbTmsorderQueryResponse struct {
    XMLName xml.Name `xml:"wlb_tmsorder_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 物流订单运单信息列表
    
    TmsOrderList   []WlbTmsOrder `json:"tms_order_list,omitempty" xml:"tms_order_list>wlb_tms_order,omitempty"`
    
    
    // 结果总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
