package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝确认收货 API返回值 
tmall.exchange.returngoods.refuse

卖家拒绝买家换货申请
*/
type TmallExchangeReturngoodsRefuseAPIResponse struct {
    model.CommonResponse
    TmallExchangeReturngoodsRefuseAPIResponseModel
}

// 卖家拒绝确认收货 成功返回结果
type TmallExchangeReturngoodsRefuseAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_exchange_returngoods_refuse_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
