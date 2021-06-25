package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号获取单个出库单发货信息 APIResponse
taobao.wlb.wms.stock.out.bill.get

通过订单号获取单个出库单发货信息
*/
type TaobaoWlbWmsStockOutBillGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsStockOutBillGetResponse `json:"taobao_wlb_wms_stock_out_bill_get_response,omitempty"`
}

type TaobaoWlbWmsStockOutBillGetResponse struct {

    // 出库信息
    StockOutInfo   *CainiaoStockOutBillStockoutinfo `json:"stock_out_info,omitempty"`

}
