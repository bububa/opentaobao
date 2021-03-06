package lsttrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderGetAPIResponse 零售通交易订单查询--品牌商视角 API返回值
// alibaba.lst.trade.order.get
//
// 根据订单id查询零售通交易订单
type AlibabaLstTradeOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeOrderGetAPIResponseModel
}

// AlibabaLstTradeOrderGetAPIResponseModel is 零售通交易订单查询--品牌商视角 成功返回结果
type AlibabaLstTradeOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaLstTradeOrderGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
