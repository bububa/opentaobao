package lstfundbill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerorderfundbilllistqueryAPIResponse 结算明细数据查询（卖家视角） API返回值
// alibaba.lst.trade.seller.order.fundbill.list.query
//
// 提供For供应商的结算接口，以交易账单维度提供开放数据，满足供应商自动化结算的诉求
type AlibabalsttradesellerorderfundbilllistqueryAPIResponse struct {
	model.CommonResponse
	AlibabalsttradesellerorderfundbilllistqueryAPIResponseModel
}

// AlibabalsttradesellerorderfundbilllistqueryAPIResponseModel is 结算明细数据查询（卖家视角） 成功返回结果
type AlibabalsttradesellerorderfundbilllistqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_fundbill_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包装类
	Result *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
