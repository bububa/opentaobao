package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单据挂起（恢复）接口 API返回值 
taobao.qimen.order.pending

ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
*/
type TaobaoQimenOrderPendingAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderPendingAPIResponseModel
}

// 单据挂起（恢复）接口 成功返回结果
type TaobaoQimenOrderPendingAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_order_pending_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *OrderPendingResponse `json:"response,omitempty" xml:"response,omitempty"`
}
