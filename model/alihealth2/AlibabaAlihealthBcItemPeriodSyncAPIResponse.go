package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbcitemperiodsyncAPIResponse 代销品效期同步 API返回值
// alibaba.alihealth.bc.item.period.sync
//
// 代销品效期同步
type AlibabaalihealthbcitemperiodsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthbcitemperiodsyncAPIResponseModel
}

// AlibabaalihealthbcitemperiodsyncAPIResponseModel is 代销品效期同步 成功返回结果
type AlibabaalihealthbcitemperiodsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_bc_item_period_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 链路跟踪id, 用于排查问题
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误编码, 编码以&#34;ALIHEALTH.BC.5&#34;开头时表示当前品仓的操作处于未知状态, 后继可通过同样的请求流水号来获取最终结果
	ErrorNo string `json:"error_no,omitempty" xml:"error_no,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求是否可重试
	CanRetry bool `json:"can_retry,omitempty" xml:"can_retry,omitempty"`
	// 当前请求是否成功-所有明细都成功才算成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
