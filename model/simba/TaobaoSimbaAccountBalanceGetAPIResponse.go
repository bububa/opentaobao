package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAccountBalanceGetAPIResponse
获取实时余额，”元”为单位 API返回值
taobao.simba.account.balance.get

获取实时余额，”元”为单位 */
type TaobaoSimbaAccountBalanceGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAccountBalanceGetAPIResponseModel
}

// TaobaoSimbaAccountBalanceGetAPIResponseModel is 获取实时余额，”元”为单位 成功返回结果
type TaobaoSimbaAccountBalanceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_account_balance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 实时余额，”元”为单位
	Balance string `json:"balance,omitempty" xml:"balance,omitempty"`
}
