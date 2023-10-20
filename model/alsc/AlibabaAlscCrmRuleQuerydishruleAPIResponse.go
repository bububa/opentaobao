package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerydishruleAPIResponse 查询品牌下的入会菜品规则 API返回值
// alibaba.alsc.crm.rule.querydishrule
//
// 查询品牌下的入会菜品规则
type AlibabaalsccrmrulequerydishruleAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmrulequerydishruleAPIResponseModel
}

// AlibabaalsccrmrulequerydishruleAPIResponseModel is 查询品牌下的入会菜品规则 成功返回结果
type AlibabaalsccrmrulequerydishruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querydishrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
