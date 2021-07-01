package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsStockOutBillGetAPIResponse
通过订单号获取单个出库单发货信息 API返回值
taobao.wlb.wms.stock.out.bill.get

通过订单号获取单个出库单发货信息 */
type TaobaoWlbWmsStockOutBillGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsStockOutBillGetAPIResponseModel
}

// TaobaoWlbWmsStockOutBillGetAPIResponseModel is 通过订单号获取单个出库单发货信息 成功返回结果
type TaobaoWlbWmsStockOutBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_stock_out_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出库信息
	StockOutInfo *CainiaoStockOutBillStockoutinfo `json:"stock_out_info,omitempty" xml:"stock_out_info,omitempty"`
}
