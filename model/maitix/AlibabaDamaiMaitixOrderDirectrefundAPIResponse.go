package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderDirectrefundAPIResponse
大麦-直接退票 API返回值
alibaba.damai.maitix.order.directrefund

大麦-退票 */
type AlibabaDamaiMaitixOrderDirectrefundAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOrderDirectrefundAPIResponseModel
}

// AlibabaDamaiMaitixOrderDirectrefundAPIResponseModel is 大麦-直接退票 成功返回结果
type AlibabaDamaiMaitixOrderDirectrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_directrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退票返回对象
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}
