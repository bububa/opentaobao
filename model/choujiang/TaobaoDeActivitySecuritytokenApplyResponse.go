package choujiang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
安全token获取 APIResponse
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
*/
type TaobaoDeActivitySecuritytokenApplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"de_activity_securitytoken_apply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功标志位
    
    Result   bool `json:"result,omitempty" xml:"