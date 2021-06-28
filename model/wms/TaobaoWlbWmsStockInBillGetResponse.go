package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取入库单信息 APIResponse
taobao.wlb.wms.stock.in.bill.get

获取入库单信息
*/
type TaobaoWlbWmsStockInBillGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_stock_in_bill_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 入库单信息
    
    StockInInfo   *CainiaoStockInBillStockininfo `json:"stock_in_info,omitempty" xml:"