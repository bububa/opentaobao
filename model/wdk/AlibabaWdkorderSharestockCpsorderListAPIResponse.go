package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockCpsorderListAPIResponse cps正向分销订单批量回流 API返回值
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
type AlibabaWdkorderSharestockCpsorderListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockCpsorderListAPIResponseModel
}

// AlibabaWdkorderSharestockCpsorderListAPIResponseModel is cps正向分销订单批量回流 成功返回结果
type AlibabaWdkorderSharestockCpsorderListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_cpsorder_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	ApiResult *AlibabaWdkorderSharestockCpsorderListApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
