package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 API返回值 
taobao.qimen.order.cancel

ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
*/
type TaobaoQimenOrderCancelAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderCancelAPIResponseModel
}

// 单据取消接口 成功返回结果
type TaobaoQimenOrderCancelAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_order_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *OrderCancelResponse `json:"response,omitempty" xml:"response,omitempty"`
}
