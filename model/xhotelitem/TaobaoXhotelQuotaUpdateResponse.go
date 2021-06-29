package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存更新接口 APIResponse
taobao.xhotel.quota.update

库存更新接口
*/
type TaobaoXhotelQuotaUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelQuotaUpdateResponse
}

type TaobaoXhotelQuotaUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_quota_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新失败补充描述消息
    
    WarnMessage   string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`

    
    // errorCode
    
    BizErrorCode   string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`

    
    // 更新失败错误信息
    
    BizErrorMsg   string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`

    
}
