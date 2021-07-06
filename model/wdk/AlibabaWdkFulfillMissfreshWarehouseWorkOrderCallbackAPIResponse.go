package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponse 每日优鲜仓作业单回传接口 API返回值
// alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback
//
// 家乐福仓作业单回传接口
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponseModel
}

// AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponseModel is 每日优鲜仓作业单回传接口 成功返回结果
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_missfresh_warehouse_work_order_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应提示信息
	RespMessage string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
	// 响应code
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
