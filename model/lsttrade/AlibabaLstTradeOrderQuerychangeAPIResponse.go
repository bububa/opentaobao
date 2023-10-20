package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderQuerychangeAPIResponse 订单id批量查询（品牌商视角） API返回值
// alibaba.lst.trade.order.querychange
//
// 根据品牌和时间段查询有变更记录的订单id
type AlibabaLstTradeOrderQuerychangeAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeOrderQuerychangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeOrderQuerychangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeOrderQuerychangeAPIResponseModel).Reset()
}

// AlibabaLstTradeOrderQuerychangeAPIResponseModel is 订单id批量查询（品牌商视角） 成功返回结果
type AlibabaLstTradeOrderQuerychangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_order_querychange_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeOrderQuerychangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstTradeOrderQuerychangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeOrderQuerychangeAPIResponse)
	},
}

// GetAlibabaLstTradeOrderQuerychangeAPIResponse 从 sync.Pool 获取 AlibabaLstTradeOrderQuerychangeAPIResponse
func GetAlibabaLstTradeOrderQuerychangeAPIResponse() *AlibabaLstTradeOrderQuerychangeAPIResponse {
	return poolAlibabaLstTradeOrderQuerychangeAPIResponse.Get().(*AlibabaLstTradeOrderQuerychangeAPIResponse)
}

// ReleaseAlibabaLstTradeOrderQuerychangeAPIResponse 将 AlibabaLstTradeOrderQuerychangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeOrderQuerychangeAPIResponse(v *AlibabaLstTradeOrderQuerychangeAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeOrderQuerychangeAPIResponse.Put(v)
}
