package yunosad

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个创意审核状态 APIResponse
yunos.ad.audit.creative.get

获取单个创意审核状态
*/
type YunosAdAuditCreativeGetAPIResponse struct {
    model.CommonResponse
    YunosAdAuditCreativeGetResponse
}

type YunosAdAuditCreativeGetResponse struct {
    XMLName xml.Name `xml:"yunos_ad_audit_creative_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 状态
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`

    
    // 审核结果
    
    Result   *CreativeAuditDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
