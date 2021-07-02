package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderListAPIResponse 五道口订单拉取 API返回值
// alibaba.wdk.order.list
//
// 五道口交易订单拉取接口
type AlibabaWdkOrderListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderListAPIResponseModel
}

// AlibabaWdkOrderListAPIResponseModel is 五道口订单拉取 成功返回结果
type AlibabaWdkOrderListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *AlibabaWdkOrderListResult `json:"result,omitempty" xml:"result,omitempty"`
}
