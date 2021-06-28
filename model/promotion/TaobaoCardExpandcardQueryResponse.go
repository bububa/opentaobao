package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
购物金卡查询 APIResponse
taobao.card.expandcard.query

购物金充值信息查询接口，会返回余额等信息。
*/
type TaobaoCardExpandcardQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"card_expandcard_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoCardExpandcardQueryResult `json:"result,omitempty" xml:"