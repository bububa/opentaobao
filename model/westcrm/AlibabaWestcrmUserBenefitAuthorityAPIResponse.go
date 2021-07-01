package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmUserBenefitAuthorityAPIResponse
获取指定用户是否含有会员权益配置菜单权限 API返回值
alibaba.westcrm.user.benefit.authority

获取指定用户是否含有会员权益配置菜单权限 */
type AlibabaWestcrmUserBenefitAuthorityAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmUserBenefitAuthorityAPIResponseModel
}

// AlibabaWestcrmUserBenefitAuthorityAPIResponseModel is 获取指定用户是否含有会员权益配置菜单权限 成功返回结果
type AlibabaWestcrmUserBenefitAuthorityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_user_benefit_authority_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
