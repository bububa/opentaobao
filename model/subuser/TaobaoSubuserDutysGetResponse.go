package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的所有职务信息列表 APIResponse
taobao.subuser.dutys.get

通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
*/
type TaobaoSubuserDutysGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubuserDutysGetResponse `json:"subuser_dutys_get_response,omitempty"` 
    TaobaoSubuserDutysGetResponse
}

/* model for simplify = false
type TaobaoSubuserDutysGetResponse struct {

    // 职务信息
    
    Dutys  struct {
        Duty  []Duty `json:"duty,omitempty"`
    } `json:"dutys,omitempty"`
    

}
*/

type TaobaoSubuserDutysGetResponse struct {

    // 职务信息
    Dutys   []Duty `json:"dutys,omitempty"`

}
