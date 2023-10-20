package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoyphordersgetAPIResponse 批量查询一盘货采购单信息 API返回值
// taobao.fenxiao.yph.orders.get
//
// 一盘货商家批量查询采购单信息
type TaobaofenxiaoyphordersgetAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaoyphordersgetAPIResponseModel
}

// TaobaofenxiaoyphordersgetAPIResponseModel is 批量查询一盘货采购单信息 成功返回结果
type TaobaofenxiaoyphordersgetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_yph_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 采购单及子采购单信息
	PurchaseOrders []OrderList `json:"purchase_orders,omitempty" xml:"purchase_orders>order_list,omitempty"`
	// 错误说明
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 接口操作时间
	OptTime string `json:"opt_time,omitempty" xml:"opt_time,omitempty"`
	// 接口返回错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 采购单查询总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
