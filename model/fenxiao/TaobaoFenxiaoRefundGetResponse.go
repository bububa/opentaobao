package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单退款信息 APIResponse
taobao.fenxiao.refund.get

分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
*/
type TaobaoFenxiaoRefundGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_refund_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 退款详情
    
    RefundDetail   *RefundDetail `json:"refund_detail,omitempty" xml:"