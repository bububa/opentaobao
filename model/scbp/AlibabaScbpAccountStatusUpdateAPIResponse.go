package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountstatusupdateAPIResponse 修改账户级别关键词推广状态 API返回值
// alibaba.scbp.account.status.update
//
// 修改账户级别关键词推广状态
type AlibabascbpaccountstatusupdateAPIResponse struct {
	model.CommonResponse
	AlibabascbpaccountstatusupdateAPIResponseModel
}

// AlibabascbpaccountstatusupdateAPIResponseModel is 修改账户级别关键词推广状态 成功返回结果
type AlibabascbpaccountstatusupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
