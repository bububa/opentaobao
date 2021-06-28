package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取DSP用户的创意审核结果 APIResponse
taobao.tanx.creatives.get

批量获取DSP用户的创意审核结果
*/
type TaobaoTanxCreativesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_creatives_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty" xml:"