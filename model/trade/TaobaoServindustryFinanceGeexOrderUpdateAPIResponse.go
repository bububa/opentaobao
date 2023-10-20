package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceGeexOrderUpdateAPIResponse 即科订单结果更新回调 API返回值
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
type TaobaoServindustryFinanceGeexOrderUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoServindustryFinanceGeexOrderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceGeexOrderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoServindustryFinanceGeexOrderUpdateAPIResponseModel).Reset()
}

// TaobaoServindustryFinanceGeexOrderUpdateAPIResponseModel is 即科订单结果更新回调 成功返回结果
type TaobaoServindustryFinanceGeexOrderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_geex_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RetryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceGeexOrderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoServindustryFinanceGeexOrderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoServindustryFinanceGeexOrderUpdateAPIResponse)
	},
}

// GetTaobaoServindustryFinanceGeexOrderUpdateAPIResponse 从 sync.Pool 获取 TaobaoServindustryFinanceGeexOrderUpdateAPIResponse
func GetTaobaoServindustryFinanceGeexOrderUpdateAPIResponse() *TaobaoServindustryFinanceGeexOrderUpdateAPIResponse {
	return poolTaobaoServindustryFinanceGeexOrderUpdateAPIResponse.Get().(*TaobaoServindustryFinanceGeexOrderUpdateAPIResponse)
}

// ReleaseTaobaoServindustryFinanceGeexOrderUpdateAPIResponse 将 TaobaoServindustryFinanceGeexOrderUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoServindustryFinanceGeexOrderUpdateAPIResponse(v *TaobaoServindustryFinanceGeexOrderUpdateAPIResponse) {
	v.Reset()
	poolTaobaoServindustryFinanceGeexOrderUpdateAPIResponse.Put(v)
}
