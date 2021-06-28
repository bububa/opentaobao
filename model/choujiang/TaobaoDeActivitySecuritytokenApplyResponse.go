package choujiang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
安全token获取 APIResponse
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
*/
type TaobaoDeActivitySecuritytokenApplyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoDeActivitySecuritytokenApplyResponse `json:"de_activity_securitytoken_apply_response,omitempty"` 
    TaobaoDeActivitySecuritytokenApplyResponse
}

/* model for simplify = false
type TaobaoDeActivitySecuritytokenApplyResponse struct {

    // 成功标志位
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoDeActivitySecuritytokenApplyResponse struct {

    // 成功标志位
    Result   bool `json:"result,omitempty"`

}
