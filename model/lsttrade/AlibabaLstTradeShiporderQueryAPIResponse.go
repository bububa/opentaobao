package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeShiporderQueryAPIResponse 供应商数据开放--发货单接口 API返回值
// alibaba.lst.trade.shiporder.query
//
// 供应商数据开放--发货单接口
type AlibabaLstTradeShiporderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeShiporderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeShiporderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeShiporderQueryAPIResponseModel).Reset()
}

// AlibabaLstTradeShiporderQueryAPIResponseModel is 供应商数据开放--发货单接口 成功返回结果
type AlibabaLstTradeShiporderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_shiporder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeShiporderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstTradeShiporderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeShiporderQueryAPIResponse)
	},
}

// GetAlibabaLstTradeShiporderQueryAPIResponse 从 sync.Pool 获取 AlibabaLstTradeShiporderQueryAPIResponse
func GetAlibabaLstTradeShiporderQueryAPIResponse() *AlibabaLstTradeShiporderQueryAPIResponse {
	return poolAlibabaLstTradeShiporderQueryAPIResponse.Get().(*AlibabaLstTradeShiporderQueryAPIResponse)
}

// ReleaseAlibabaLstTradeShiporderQueryAPIResponse 将 AlibabaLstTradeShiporderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeShiporderQueryAPIResponse(v *AlibabaLstTradeShiporderQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeShiporderQueryAPIResponse.Put(v)
}
