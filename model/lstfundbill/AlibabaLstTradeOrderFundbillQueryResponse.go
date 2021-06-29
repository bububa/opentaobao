package lstfundbill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
结算明细数据查询（品牌商视角） API返回值 
alibaba.lst.trade.order.fundbill.query

按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。
*/
type AlibabaLstTradeOrderFundbillQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeOrderFundbillQueryResponse
}

// 结算明细数据查询（品牌商视角） 成功返回结果
type AlibabaLstTradeOrderFundbillQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_order_fundbill_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 包装类
    Result   *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
