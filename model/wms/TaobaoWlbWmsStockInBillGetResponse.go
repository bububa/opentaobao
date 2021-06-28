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
    TaobaoWlbWmsStockInBillGetResponse
}

type TaobaoWlbWmsStockInBillGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_stock_in_bill_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 入库单信息
    
    StockInInfo   *CainiaoStockInBillStockininfo `json:"stock_in_info,omitempty" xml:"stock_in_info,omitempty"`

    
}
