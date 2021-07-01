package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝订单已发货通知接口 API返回值 
taobao.wlb.order.consign

如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
*/
type TaobaoWlbOrderConsignAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderConsignAPIResponseModel
}

// 物流宝订单已发货通知接口 成功返回结果
type TaobaoWlbOrderConsignAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_order_consign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改时间
    ModifyTime   string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
}
