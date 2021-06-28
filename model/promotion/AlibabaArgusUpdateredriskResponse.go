package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新红线价格 APIResponse
alibaba.argus.updateredrisk

商品健康中心新增红线价格规则
*/
type AlibabaArgusUpdateredriskAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_argus_updateredrisk_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"