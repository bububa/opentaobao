package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务订单详情 APIResponse
taobao.xhotel.data.service.order.detail

服务订单详情top接口构建
*/
type TaobaoXhotelDataServiceOrderDetailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_data_service_order_detail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoXhotelDataServiceOrderDetailResultSet `json:"result,omitempty" xml:"