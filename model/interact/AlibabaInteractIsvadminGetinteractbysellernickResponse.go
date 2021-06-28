package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据sellerNick获取互动实例列表 APIResponse
alibaba.interact.isvadmin.getinteractbysellernick

根据sellerNick获取互动实例列表
*/
type AlibabaInteractIsvadminGetinteractbysellernickAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_isvadmin_getinteractbysellernick_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"