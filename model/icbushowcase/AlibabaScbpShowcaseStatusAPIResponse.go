package icbushowcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpShowcaseStatusAPIResponse
橱窗状态 API返回值
alibaba.scbp.showcase.status

查询橱窗状态，如总数、可用数量 */
type AlibabaScbpShowcaseStatusAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseStatusAPIResponseModel
}

// AlibabaScbpShowcaseStatusAPIResponseModel is 橱窗状态 成功返回结果
type AlibabaScbpShowcaseStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 全部橱窗数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前已用的橱窗数
	CurrentCount int64 `json:"current_count,omitempty" xml:"current_count,omitempty"`
}
