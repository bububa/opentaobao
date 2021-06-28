package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫互动奖池列表 APIResponse
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表
*/
type AlibabaInteractIsvadminAllpondsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_isvadmin_allponds_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"