package lsttrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeShiporderQueryAPIResponse
供应商数据开放--发货单接口 API返回值
alibaba.lst.trade.shiporder.query

供应商数据开放--发货单接口 */
type AlibabaLstTradeShiporderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeShiporderQueryAPIResponseModel
}

// AlibabaLstTradeShiporderQueryAPIResponseModel is 供应商数据开放--发货单接口 成功返回结果
type AlibabaLstTradeShiporderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_shiporder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
