package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据互动实例ID查询奖池信息 APIResponse
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息
*/
type AlibabaInteractIsvadminGetpondbyinteractAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_isvadmin_getpondbyinteract_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 奖池信息
    
    Data   *PrizePondDTO `json:"data,omitempty" xml:"