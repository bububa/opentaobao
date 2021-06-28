package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘定向优惠打标接口 APIResponse
taobao.ump.shoutaotag.add

手淘定向优惠的优惠标签打标接口
给特定的手淘买家打上优惠标记，标记承载在自己的业务标签库中，标签有效期为7天。
*/
type TaobaoUmpShoutaotagAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpShoutaotagAddResponse `json:"ump_shoutaotag_add_response,omitempty"` 
    TaobaoUmpShoutaotagAddResponse
}

/* model for simplify = false
type TaobaoUmpShoutaotagAddResponse struct {

    // 是否打标成功
    
    AddResult   bool `json:"add_result,omitempty"`
    

}
*/

type TaobaoUmpShoutaotagAddResponse struct {

    // 是否打标成功
    AddResult   bool `json:"add_result,omitempty"`

}
