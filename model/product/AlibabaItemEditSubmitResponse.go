package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑提交schema信息 APIResponse
alibaba.item.edit.submit

商品编辑提交schema信息
*/
type AlibabaItemEditSubmitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaItemEditSubmitResponse `json:"alibaba_item_edit_submit_response,omitempty"` 
    AlibabaItemEditSubmitResponse
}

/* model for simplify = false
type AlibabaItemEditSubmitResponse struct {

    // 商品更新时间
    
    UpdateTime   string `json:"update_time,omitempty"`
    

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 商品所属市场
    
    Market   string `json:"market,omitempty"`
    

}
*/

type AlibabaItemEditSubmitResponse struct {

    // 商品更新时间
    UpdateTime   string `json:"update_time,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品所属市场
    Market   string `json:"market,omitempty"`

}
