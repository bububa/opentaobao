package yunosad

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个创意预审接口 API返回值 
yunos.ad.audit.creative.add

YunOS广告业务第三方DSP单个创意预审接口
*/
type YunosAdAuditCreativeAddAPIResponse struct {
    model.CommonResponse
    YunosAdAuditCreativeAddResponse
}

// 单个创意预审接口 成功返回结果
type YunosAdAuditCreativeAddResponse struct {
    XMLName xml.Name `xml:"yunos_ad_audit_creative_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // statusCode
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // isOk
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
}
