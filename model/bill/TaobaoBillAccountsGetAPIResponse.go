package bill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBillAccountsGetAPIResponse 查询费用科目信息(限自研商家) API返回值
// taobao.bill.accounts.get
//
// 查询费用账户信息
type TaobaoBillAccountsGetAPIResponse struct {
	model.CommonResponse
	TaobaoBillAccountsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBillAccountsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBillAccountsGetAPIResponseModel).Reset()
}

// TaobaoBillAccountsGetAPIResponseModel is 查询费用科目信息(限自研商家) 成功返回结果
type TaobaoBillAccountsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bill_accounts_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的科目信息
	Accounts []Account `json:"accounts,omitempty" xml:"accounts>account,omitempty"`
	// 返回记录行数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBillAccountsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Accounts = m.Accounts[:0]
	m.TotalResults = 0
}

var poolTaobaoBillAccountsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBillAccountsGetAPIResponse)
	},
}

// GetTaobaoBillAccountsGetAPIResponse 从 sync.Pool 获取 TaobaoBillAccountsGetAPIResponse
func GetTaobaoBillAccountsGetAPIResponse() *TaobaoBillAccountsGetAPIResponse {
	return poolTaobaoBillAccountsGetAPIResponse.Get().(*TaobaoBillAccountsGetAPIResponse)
}

// ReleaseTaobaoBillAccountsGetAPIResponse 将 TaobaoBillAccountsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBillAccountsGetAPIResponse(v *TaobaoBillAccountsGetAPIResponse) {
	v.Reset()
	poolTaobaoBillAccountsGetAPIResponse.Put(v)
}
