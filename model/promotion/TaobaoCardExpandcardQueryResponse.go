package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
购物金卡查询 APIResponse
taobao.card.expandcard.query

购物金充值信息查询接口，会返回余额等信息。
*/
type TaobaoCardExpandcardQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCardExpandcardQueryResponse `json:"card_expandcard_query_response,omitempty"` 
    TaobaoCardExpandcardQueryResponse
}

/* model for simplify = false
type TaobaoCardExpandcardQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoCardExpandcardQueryResult  *TaobaoCardExpandcardQueryResult `json:"taobao_card_expandcard_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoCardExpandcardQueryResponse struct {

    // 接口返回model
    Result   *TaobaoCardExpandcardQueryResult `json:"result,omitempty"`

}
