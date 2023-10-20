package aliexpress

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsAbnormalorderQueryAPIResponse 异常订单查询 API返回值
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
type AliexpressLogisticsAbnormalorderQueryAPIResponse struct {
	model.CommonResponse
	AliexpressLogisticsAbnormalorderQueryAPIResponseModel
}

// AliexpressLogisticsAbnormalorderQueryAPIResponseModel is 异常订单查询 成功返回结果
type AliexpressLogisticsAbnormalorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_logistics_abnormalorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单详情
	Result []AeopWarehouseAbnormalOrderResult `json:"result,omitempty" xml:"result>aeop_warehouse_abnormal_order_result,omitempty"`
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 是否成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
