package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletcheckbalanceAPIResponse 商家预存余额检查 API返回值
// alibaba.aliqin.flow.wallet.check.balance
//
// 检查商家CRM预存余额是否足够进行活动
type AlibabaaliqinflowwalletcheckbalanceAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinflowwalletcheckbalanceAPIResponseModel
}

// AlibabaaliqinflowwalletcheckbalanceAPIResponseModel is 商家预存余额检查 成功返回结果
type AlibabaaliqinflowwalletcheckbalanceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_check_balance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 余额是否大于校验值，大于返回true，小于返回false
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
