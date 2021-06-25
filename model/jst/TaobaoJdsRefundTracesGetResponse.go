package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单条退款跟踪详情 APIResponse
taobao.jds.refund.traces.get

获取聚石塔数据共享的交易全链路的退款信息
*/
type TaobaoJdsRefundTracesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJdsRefundTracesGetResponse `json:"taobao_jds_refund_traces_get_response,omitempty"`
}

type TaobaoJdsRefundTracesGetResponse struct {

    // 退款跟踪列表
    Traces   []RefundTrace `json:"traces,omitempty"`

    // 用户在全链路系统中的状态(比如是否存在)
    UserStatus   string `json:"user_status,omitempty"`

}
