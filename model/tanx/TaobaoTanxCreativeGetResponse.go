package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审核创意状态 APIResponse
taobao.tanx.creative.get

获取单个审核创意状态
*/
type TaobaoTanxCreativeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_creative_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty" xml:"