package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowAccountGetAPIResponse 获取信息流账户详情 API返回值
// taobao.feedflow.account.get
//
// 获取账户信息接口。
// (1) BP显示余额 (字段 ：banlance ) = 现金余额(字段：cash_balance) + 赠款余额；
// (2) 可用余额(字段：availableBalance) = BP显示余额
// (3) 红包(字段：redPacket)
type TaobaoFeedflowAccountGetAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowAccountGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowAccountGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowAccountGetAPIResponseModel).Reset()
}

// TaobaoFeedflowAccountGetAPIResponseModel is 获取信息流账户详情 成功返回结果
type TaobaoFeedflowAccountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_account_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowAccountGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowAccountGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowAccountGetAPIResponse)
	},
}

// GetTaobaoFeedflowAccountGetAPIResponse 从 sync.Pool 获取 TaobaoFeedflowAccountGetAPIResponse
func GetTaobaoFeedflowAccountGetAPIResponse() *TaobaoFeedflowAccountGetAPIResponse {
	return poolTaobaoFeedflowAccountGetAPIResponse.Get().(*TaobaoFeedflowAccountGetAPIResponse)
}

// ReleaseTaobaoFeedflowAccountGetAPIResponse 将 TaobaoFeedflowAccountGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowAccountGetAPIResponse(v *TaobaoFeedflowAccountGetAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowAccountGetAPIResponse.Put(v)
}
