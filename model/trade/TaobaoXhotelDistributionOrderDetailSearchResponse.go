package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道订单详情查询 APIResponse
taobao.xhotel.distribution.order.detail.search

该接口用于提供酒店分销渠道的订单详情查询
*/
type TaobaoXhotelDistributionOrderDetailSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_distribution_order_detail_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    Error   string `json:"error,omitempty" xml:"