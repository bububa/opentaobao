package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
应用上下架操作 APIResponse
yunos.tvpubadmin.content.app.onoffappbylicense

应用上下架操作
*/
type YunosTvpubadminContentAppOnoffappbylicenseAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAppOnoffappbylicenseResponse
}

type YunosTvpubadminContentAppOnoffappbylicenseResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_onoffappbylicense_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true/false
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
