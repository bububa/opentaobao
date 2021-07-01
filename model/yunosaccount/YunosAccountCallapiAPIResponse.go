package yunosaccount

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosAccountCallapiAPIResponse
调用YunOS账号开放API API返回值
yunos.account.callapi

YunOS账号客户端对外开放的api通过top暴露 */
type YunosAccountCallapiAPIResponse struct {
	model.CommonResponse
	YunosAccountCallapiAPIResponseModel
}

// YunosAccountCallapiAPIResponseModel is 调用YunOS账号开放API 成功返回结果
type YunosAccountCallapiAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_account_callapi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
