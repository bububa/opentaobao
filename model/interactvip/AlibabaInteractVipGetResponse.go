package interactvip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员淘气值获取 APIResponse
alibaba.interact.vip.get

提供用户淘气值&用户角色身份查询
*/
type AlibabaInteractVipGetAPIResponse struct {
    model.CommonResponse
    AlibabaInteractVipGetResponse
}

type AlibabaInteractVipGetResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_vip_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
