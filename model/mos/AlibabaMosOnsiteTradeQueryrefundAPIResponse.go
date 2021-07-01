package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款查询 API返回值 
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。
*/
type AlibabaMosOnsiteTradeQueryrefundAPIResponse struct {
    model.CommonResponse
    AlibabaMosOnsiteTradeQueryrefundAPIResponseModel
}

// 退款查询 成功返回结果
type AlibabaMosOnsiteTradeQueryrefundAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mos_onsite_trade_queryrefund_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaMosOnsiteTradeQueryrefundResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
