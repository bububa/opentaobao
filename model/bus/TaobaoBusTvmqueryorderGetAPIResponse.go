package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmqueryordergetAPIResponse 线下自助机查询订单信息 API返回值
// taobao.bus.tvmqueryorder.get
//
// 查询订单详情
type TaobaobustvmqueryordergetAPIResponse struct {
	model.CommonResponse
	TaobaobustvmqueryordergetAPIResponseModel
}

// TaobaobustvmqueryordergetAPIResponseModel is 线下自助机查询订单信息 成功返回结果
type TaobaobustvmqueryordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmqueryorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// tvmBusOrderLineInfo
	TvmBusOrderLineInfo *TvmBusOrderLineInfo `json:"tvm_bus_order_line_info,omitempty" xml:"tvm_bus_order_line_info,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
