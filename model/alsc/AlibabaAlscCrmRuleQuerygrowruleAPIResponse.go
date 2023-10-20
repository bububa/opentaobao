package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerygrowruleAPIResponse 查询品牌下的会员成长规则 API返回值
// alibaba.alsc.crm.rule.querygrowrule
//
// 查询品牌下的会员成长规则
type AlibabaalsccrmrulequerygrowruleAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmrulequerygrowruleAPIResponseModel
}

// AlibabaalsccrmrulequerygrowruleAPIResponseModel is 查询品牌下的会员成长规则 成功返回结果
type AlibabaalsccrmrulequerygrowruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querygrowrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
