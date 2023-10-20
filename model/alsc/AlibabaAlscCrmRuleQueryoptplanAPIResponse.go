package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequeryoptplanAPIResponse 查询运营计划 API返回值
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
type AlibabaalsccrmrulequeryoptplanAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmrulequeryoptplanAPIResponseModel
}

// AlibabaalsccrmrulequeryoptplanAPIResponseModel is 查询运营计划 成功返回结果
type AlibabaalsccrmrulequeryoptplanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_queryoptplan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求参数
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
