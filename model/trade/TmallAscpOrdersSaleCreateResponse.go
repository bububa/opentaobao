package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ASCP渠道中心销售单创建接口 APIResponse
tmall.ascp.orders.sale.create

ASCP渠道中心销售单创建接口
*/
type TmallAscpOrdersSaleCreateAPIResponse struct {
    model.CommonResponse
    Response *TmallAscpOrdersSaleCreateResponse `json:"tmall_ascp_orders_sale_create_response,omitempty"`
}

type TmallAscpOrdersSaleCreateResponse struct {

    // result
    Result   *TmallAscpOrdersSaleCreateResultDo `json:"result,omitempty"`

}
