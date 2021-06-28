package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品变更消息 APIResponse
cainiao.cntec.item.change.message

供货商商品信息变更消息
*/
type CainiaoCntecItemChangeMessageAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCntecItemChangeMessageResponse `json:"cainiao_cntec_item_change_message_response,omitempty"` 
    CainiaoCntecItemChangeMessageResponse
}

/* model for simplify = false
type CainiaoCntecItemChangeMessageResponse struct {

    // 调用返回的result结构体
    
    Result  *struct {
        CainiaoCntecItemChangeMessageResult  *CainiaoCntecItemChangeMessageResult `json:"cainiao_cntec_item_change_message_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCntecItemChangeMessageResponse struct {

    // 调用返回的result结构体
    Result   *CainiaoCntecItemChangeMessageResult `json:"result,omitempty"`

}
