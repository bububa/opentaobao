package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkRefundAggregateAPIResponse
淘鲜达退款单按门店聚合查询 API返回值
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询 */
type AlibabaWdkRefundAggregateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkRefundAggregateAPIResponseModel
}

// AlibabaWdkRefundAggregateAPIResponseModel is 淘鲜达退款单按门店聚合查询 成功返回结果
type AlibabaWdkRefundAggregateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_refund_aggregate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundAggregateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
