package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取入库单信息 APIResponse
taobao.wlb.wms.stock.in.bill.get

获取入库单信息
*/
type TaobaoWlbWmsStockInBillGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsStockInBillGetResponse `json:"taobao_wlb_wms_stock_in_bill_get_response,omitempty"`
}

type TaobaoWlbWmsStockInBillGetResponse struct {

    // 入库单信息
    StockInInfo   *CainiaoStockInBillStockininfo `json:"stock_in_info,omitempty"`

}
