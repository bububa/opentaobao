package filmtfavatar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院卖品账单--退款账单 APIResponse
taobao.film.tfavatar.bill.sale.refund.query

获取影院卖品账单--退款账单
*/
type TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse struct {
    model.CommonResponse
    TaobaoFilmTfavatarBillSaleRefundQueryResponse
}

type TaobaoFilmTfavatarBillSaleRefundQueryResponse struct {
    XMLName xml.Name `xml:"film_tfavatar_bill_sale_refund_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoFilmTfavatarBillSaleRefundQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
