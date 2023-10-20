package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse 北京小蜜蜂配作业回传 API返回值
// alibaba.wdk.fulfill.dms.ebeecake.work.order.callback
//
// 北京小蜜蜂配作业回传。
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponseModel
}

// AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponseModel is 北京小蜜蜂配作业回传 成功返回结果
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_dms_ebeecake_work_order_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应提示信息
	RespMessage string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
	// 响应code
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
