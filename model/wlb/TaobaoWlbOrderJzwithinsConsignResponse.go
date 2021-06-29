package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 API返回值 
taobao.wlb.order.jzwithins.consign

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
*/
type TaobaoWlbOrderJzwithinsConsignAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderJzwithinsConsignResponse
}

// 家装发货接口 成功返回结果
type TaobaoWlbOrderJzwithinsConsignResponse struct {
    XMLName xml.Name `xml:"wlb_order_jzwithins_consign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 发货成功或者失败
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 发货返回信息，如果发货错误则报出对应错误
    ResultInfo   string `json:"result_info,omitempty" xml:"result_info,omitempty"`
}
