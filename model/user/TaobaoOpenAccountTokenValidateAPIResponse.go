package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccounttokenvalidateAPIResponse open account token验证 API返回值
// taobao.open.account.token.validate
//
// open account token验证
type TaobaoopenaccounttokenvalidateAPIResponse struct {
	model.CommonResponse
	TaobaoopenaccounttokenvalidateAPIResponseModel
}

// TaobaoopenaccounttokenvalidateAPIResponseModel is open account token验证 成功返回结果
type TaobaoopenaccounttokenvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_token_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 验证成功返回token中的信息
	Data *OpenAccountTokenValidateResult `json:"data,omitempty" xml:"data,omitempty"`
}
