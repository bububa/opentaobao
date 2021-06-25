package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询订单基本信息 APIResponse
alibaba.ele.fengniao.order.query

查询订单基本信息
*/
type AlibabaEleFengniaoOrderQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoOrderQueryResponse `json:"alibaba_ele_fengniao_order_query_response,omitempty"`
}

type AlibabaEleFengniaoOrderQueryResponse struct {

    // kvs
    Kvs   []Kvs `json:"kvs,omitempty"`

    // 收件人经度
    ReceiverLongitude   string `json:"receiver_longitude,omitempty"`

    // 寄件人纬度
    TransportLatitude   string `json:"transport_latitude,omitempty"`

    // 寄件人经度
    TransportLongitude   string `json:"transport_longitude,omitempty"`

    // 订单号
    OrderId   int64 `json:"order_id,omitempty"`

    // 收件人纬度
    ReceiverLatitude   string `json:"receiver_latitude,omitempty"`

}
