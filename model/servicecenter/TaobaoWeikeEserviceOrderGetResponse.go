package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单查询 APIResponse
taobao.weike.eservice.order.get

用于客服外包中服务商查询订单列表
*/
type TaobaoWeikeEserviceOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoWeikeEserviceOrderGetResponse
}

type TaobaoWeikeEserviceOrderGetResponse struct {
    XMLName xml.Name `xml:"weike_eservice_order_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 记录总记录数
    
    TotalSize   int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`

    
    // 订单列表
    
    OrderList   []Order `json:"order_list,omitempty" xml:"order_list>order,omitempty"`
    
    
}
