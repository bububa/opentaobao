package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulelevelquerylevelruleAPIResponse 查询会员等级规则 API返回值
// alibaba.alsc.crm.rule.level.querylevelrule
//
// 查询会员等级规则
type AlibabaalsccrmrulelevelquerylevelruleAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmrulelevelquerylevelruleAPIResponseModel
}

// AlibabaalsccrmrulelevelquerylevelruleAPIResponseModel is 查询会员等级规则 成功返回结果
type AlibabaalsccrmrulelevelquerylevelruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_level_querylevelrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 会员等级规则
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
