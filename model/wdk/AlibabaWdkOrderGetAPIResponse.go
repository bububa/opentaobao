package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderGetAPIResponse 交易订单详情查询 API返回值
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
type AlibabaWdkOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderGetAPIResponseModel
}

// AlibabaWdkOrderGetAPIResponseModel is 交易订单详情查询 成功返回结果
type AlibabaWdkOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *AlibabaWdkOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
