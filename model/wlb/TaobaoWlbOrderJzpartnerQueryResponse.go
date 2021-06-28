package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询家装服务商列表 APIResponse
taobao.wlb.order.jzpartner.query

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
*/
type TaobaoWlbOrderJzpartnerQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_order_jzpartner_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口查询成功或者失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"