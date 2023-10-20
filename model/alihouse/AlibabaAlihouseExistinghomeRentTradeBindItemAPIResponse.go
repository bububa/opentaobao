package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse 交易绑定商品 API返回值
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
type AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponseModel is 交易绑定商品 成功返回结果
type AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_rent_trade_bind_item_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值对象
	Result *AlibabaAlihouseExistinghomeRentTradeBindItemResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse
func GetAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse() *AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse {
	return poolAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse.Get().(*AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse 将 AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse(v *AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse.Put(v)
}
