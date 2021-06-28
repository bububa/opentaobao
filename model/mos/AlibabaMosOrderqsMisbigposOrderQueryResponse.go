package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos新选单退 APIResponse
alibaba.mos.orderqs.misbigpos.order.query

大pos新选单退
*/
type AlibabaMosOrderqsMisbigposOrderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_orderqs_misbigpos_order_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    ErrsMsg   string `json:"errs_msg,omitempty" xml:"