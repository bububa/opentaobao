package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAccountGetBalanceAPIResponse 获取账户余额，现金余额 API返回值
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
type TaobaoUniversalbpAccountGetBalanceAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpAccountGetBalanceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAccountGetBalanceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpAccountGetBalanceAPIResponseModel).Reset()
}

// TaobaoUniversalbpAccountGetBalanceAPIResponseModel is 获取账户余额，现金余额 成功返回结果
type TaobaoUniversalbpAccountGetBalanceAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_account_get_balance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAccountGetBalanceTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAccountGetBalanceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpAccountGetBalanceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAccountGetBalanceAPIResponse)
	},
}

// GetTaobaoUniversalbpAccountGetBalanceAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpAccountGetBalanceAPIResponse
func GetTaobaoUniversalbpAccountGetBalanceAPIResponse() *TaobaoUniversalbpAccountGetBalanceAPIResponse {
	return poolTaobaoUniversalbpAccountGetBalanceAPIResponse.Get().(*TaobaoUniversalbpAccountGetBalanceAPIResponse)
}

// ReleaseTaobaoUniversalbpAccountGetBalanceAPIResponse 将 TaobaoUniversalbpAccountGetBalanceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpAccountGetBalanceAPIResponse(v *TaobaoUniversalbpAccountGetBalanceAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpAccountGetBalanceAPIResponse.Put(v)
}
