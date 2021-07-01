package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenBatchpushticketAPIResponse
大麦换验平台-第三方对外开放-票单接口batchPushTicket API返回值
alibaba.damai.mev.open.batchpushticket

批量推送票单 */
type AlibabaDamaiMevOpenBatchpushticketAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenBatchpushticketAPIResponseModel
}

// AlibabaDamaiMevOpenBatchpushticketAPIResponseModel is 大麦换验平台-第三方对外开放-票单接口batchPushTicket 成功返回结果
type AlibabaDamaiMevOpenBatchpushticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_batchpushticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenBatchpushticketResult `json:"result,omitempty" xml:"result,omitempty"`
}
