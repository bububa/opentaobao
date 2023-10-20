package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsShiftGetAPIResponse 移库单获取 API返回值
// alibaba.wdk.ums.shift.get
//
// 移库单获取
type AlibabaWdkUmsShiftGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsShiftGetAPIResponseModel
}

// AlibabaWdkUmsShiftGetAPIResponseModel is 移库单获取 成功返回结果
type AlibabaWdkUmsShiftGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_shift_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
