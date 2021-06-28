package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 APIResponse
taobao.wlb.order.jz.consign

家装类订单使用该接口发货
*/
type TaobaoWlbOrderJzConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_order_jz_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultErrorCode   string `json:"result_error_code,omitempty" xml:"