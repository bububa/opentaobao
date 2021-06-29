package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 API返回值 
taobao.film.data.third.party.refund.order

淘票票第三方退票接口
*/
type TaobaoFilmDataThirdPartyRefundOrderAPIResponse struct {
    model.CommonResponse
    TaobaoFilmDataThirdPartyRefundOrderResponse
}

// 退票接口 成功返回结果
type TaobaoFilmDataThirdPartyRefundOrderResponse struct {
    XMLName xml.Name `xml:"film_data_third_party_refund_order_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`
}
