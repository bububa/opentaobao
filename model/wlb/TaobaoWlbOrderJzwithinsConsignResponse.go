package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 APIResponse
taobao.wlb.order.jzwithins.consign

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
*/
type TaobaoWlbOrderJzwithinsConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_order_jzwithins_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 发货成功或者失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"