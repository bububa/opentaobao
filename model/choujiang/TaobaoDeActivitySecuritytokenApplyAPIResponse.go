package choujiang

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeActivitySecuritytokenApplyAPIResponse
安全token获取 API返回值
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用 */
type TaobaoDeActivitySecuritytokenApplyAPIResponse struct {
	model.CommonResponse
	TaobaoDeActivitySecuritytokenApplyAPIResponseModel
}

// TaobaoDeActivitySecuritytokenApplyAPIResponseModel is 安全token获取 成功返回结果
type TaobaoDeActivitySecuritytokenApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"de_activity_securitytoken_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功标志位
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
