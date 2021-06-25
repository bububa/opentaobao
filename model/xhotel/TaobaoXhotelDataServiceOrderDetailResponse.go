package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务订单详情 APIResponse
taobao.xhotel.data.service.order.detail

服务订单详情top接口构建
*/
type TaobaoXhotelDataServiceOrderDetailAPIResponse struct {
    model.CommonResponse
    Response *TaobaoXhotelDataServiceOrderDetailResponse `json:"taobao_xhotel_data_service_order_detail_response,omitempty"`
}

type TaobaoXhotelDataServiceOrderDetailResponse struct {

    // result
    Result   *TaobaoXhotelDataServiceOrderDetailResultSet `json:"result,omitempty"`

}
