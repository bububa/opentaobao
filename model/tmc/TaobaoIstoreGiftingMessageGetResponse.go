package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
gifting消息获取 APIResponse
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息
*/
type TaobaoIstoreGiftingMessageGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoIstoreGiftingMessageGetResponse `json:"istore_gifting_message_get_response,omitempty"` 
    TaobaoIstoreGiftingMessageGetResponse
}

/* model for simplify = false
type TaobaoIstoreGiftingMessageGetResponse struct {

    // result
    
    Result  *struct {
        TaobaoIstoreGiftingMessageGetResultDto  *TaobaoIstoreGiftingMessageGetResultDto `json:"taobao_istore_gifting_message_get_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoIstoreGiftingMessageGetResponse struct {

    // result
    Result   *TaobaoIstoreGiftingMessageGetResultDto `json:"result,omitempty"`

}
