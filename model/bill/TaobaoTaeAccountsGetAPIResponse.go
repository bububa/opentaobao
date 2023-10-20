package bill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeAccountsGetAPIResponse tae查询费用科目信息 API返回值
// taobao.tae.accounts.get
//
// tae查询费用科目信息
type TaobaoTaeAccountsGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaeAccountsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaeAccountsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaeAccountsGetAPIResponseModel).Reset()
}

// TaobaoTaeAccountsGetAPIResponseModel is tae查询费用科目信息 成功返回结果
type TaobaoTaeAccountsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tae_accounts_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的科目信息
	Accounts []TopAccountDto `json:"accounts,omitempty" xml:"accounts>top_account_dto,omitempty"`
	// 返回记录行数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaeAccountsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Accounts = m.Accounts[:0]
	m.TotalResults = 0
}

var poolTaobaoTaeAccountsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaeAccountsGetAPIResponse)
	},
}

// GetTaobaoTaeAccountsGetAPIResponse 从 sync.Pool 获取 TaobaoTaeAccountsGetAPIResponse
func GetTaobaoTaeAccountsGetAPIResponse() *TaobaoTaeAccountsGetAPIResponse {
	return poolTaobaoTaeAccountsGetAPIResponse.Get().(*TaobaoTaeAccountsGetAPIResponse)
}

// ReleaseTaobaoTaeAccountsGetAPIResponse 将 TaobaoTaeAccountsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaeAccountsGetAPIResponse(v *TaobaoTaeAccountsGetAPIResponse) {
	v.Reset()
	poolTaobaoTaeAccountsGetAPIResponse.Put(v)
}
