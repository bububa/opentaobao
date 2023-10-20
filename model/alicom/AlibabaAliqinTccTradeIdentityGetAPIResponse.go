package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTccTradeIdentityGetAPIResponse 运营商获得用户身份信息 API返回值
// alibaba.aliqin.tcc.trade.identity.get
//
// 天猫网厅运营商官方旗舰店获取用户身份信息
type AlibabaAliqinTccTradeIdentityGetAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinTccTradeIdentityGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinTccTradeIdentityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinTccTradeIdentityGetAPIResponseModel).Reset()
}

// AlibabaAliqinTccTradeIdentityGetAPIResponseModel is 运营商获得用户身份信息 成功返回结果
type AlibabaAliqinTccTradeIdentityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_tcc_trade_identity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回身份信息
	Result *IdentityInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinTccTradeIdentityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinTccTradeIdentityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTccTradeIdentityGetAPIResponse)
	},
}

// GetAlibabaAliqinTccTradeIdentityGetAPIResponse 从 sync.Pool 获取 AlibabaAliqinTccTradeIdentityGetAPIResponse
func GetAlibabaAliqinTccTradeIdentityGetAPIResponse() *AlibabaAliqinTccTradeIdentityGetAPIResponse {
	return poolAlibabaAliqinTccTradeIdentityGetAPIResponse.Get().(*AlibabaAliqinTccTradeIdentityGetAPIResponse)
}

// ReleaseAlibabaAliqinTccTradeIdentityGetAPIResponse 将 AlibabaAliqinTccTradeIdentityGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinTccTradeIdentityGetAPIResponse(v *AlibabaAliqinTccTradeIdentityGetAPIResponse) {
	v.Reset()
	poolAlibabaAliqinTccTradeIdentityGetAPIResponse.Put(v)
}
