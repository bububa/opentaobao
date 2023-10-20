package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsStockInBillGetAPIResponse 获取入库单信息 API返回值
// taobao.wlb.wms.stock.in.bill.get
//
// 获取入库单信息
type TaobaoWlbWmsStockInBillGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsStockInBillGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsStockInBillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsStockInBillGetAPIResponseModel).Reset()
}

// TaobaoWlbWmsStockInBillGetAPIResponseModel is 获取入库单信息 成功返回结果
type TaobaoWlbWmsStockInBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_stock_in_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 入库单信息
	StockInInfo *CainiaoStockInBillStockininfo `json:"stock_in_info,omitempty" xml:"stock_in_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsStockInBillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StockInInfo = nil
}

var poolTaobaoWlbWmsStockInBillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsStockInBillGetAPIResponse)
	},
}

// GetTaobaoWlbWmsStockInBillGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsStockInBillGetAPIResponse
func GetTaobaoWlbWmsStockInBillGetAPIResponse() *TaobaoWlbWmsStockInBillGetAPIResponse {
	return poolTaobaoWlbWmsStockInBillGetAPIResponse.Get().(*TaobaoWlbWmsStockInBillGetAPIResponse)
}

// ReleaseTaobaoWlbWmsStockInBillGetAPIResponse 将 TaobaoWlbWmsStockInBillGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsStockInBillGetAPIResponse(v *TaobaoWlbWmsStockInBillGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsStockInBillGetAPIResponse.Put(v)
}
