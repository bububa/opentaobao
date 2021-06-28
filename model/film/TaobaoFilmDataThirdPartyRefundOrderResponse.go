package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 APIResponse
taobao.film.data.third.party.refund.order

淘票票第三方退票接口
*/
type TaobaoFilmDataThirdPartyRefundOrderAPIResponse struct {
    model.CommonResponse
    TaobaoFilmDataThirdPartyRefundOrderResponse
}

type TaobaoFilmDataThirdPartyRefundOrderResponse struct {
    XMLName xml.Name `xml:"film_data_third_party_refund_order_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
