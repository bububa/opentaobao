package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售订单总额统计 APIResponse
taobao.daogoubao.order.statistics.total

对接千牛端数字中心
*/
type TaobaoDaogoubaoOrderStatisticsTotalAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDaogoubaoOrderStatisticsTotalResponse `json:"taobao_daogoubao_order_statistics_total_response,omitempty"`
}

type TaobaoDaogoubaoOrderStatisticsTotalResponse struct {

    // result
    Result   *OrderStatisticsResult `json:"result,omitempty"`

}
