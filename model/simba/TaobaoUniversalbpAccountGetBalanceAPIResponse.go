package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpaccountgetbalanceAPIResponse 获取账户余额，现金余额 API返回值
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
type TaobaouniversalbpaccountgetbalanceAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpaccountgetbalanceAPIResponseModel
}

// TaobaouniversalbpaccountgetbalanceAPIResponseModel is 获取账户余额，现金余额 成功返回结果
type TaobaouniversalbpaccountgetbalanceAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_account_get_balance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpaccountgetbalanceTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
