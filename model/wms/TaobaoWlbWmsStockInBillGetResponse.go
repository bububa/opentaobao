package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取入库单信息 API返回值 
taobao.wlb.wms.stock.in.bill.get

获取入库单信息
*/
type TaobaoWlbWmsStockInBillGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsStockInBillGetResponse
}

// 获取入库单信息 成功返回结果
type TaobaoWlbWmsStockInBillGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_stock_in_bill_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 入库单信息
    StockInInfo   *CainiaoStockInBillStockininfo `json:"stock_in_info,omitempty" xml:"stock_in_info,omitempty"`
}
