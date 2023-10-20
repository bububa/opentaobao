package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbilllistAPIResponse 五道口账单拉取接口 API返回值
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
type AlibabawdkbilllistAPIResponse struct {
	model.CommonResponse
	AlibabawdkbilllistAPIResponseModel
}

// AlibabawdkbilllistAPIResponseModel is 五道口账单拉取接口 成功返回结果
type AlibabawdkbilllistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bill_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口响应结果
	ApiResult *AlibabawdkbilllistApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
