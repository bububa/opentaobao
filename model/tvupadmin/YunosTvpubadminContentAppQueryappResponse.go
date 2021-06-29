package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询应用信息 APIResponse
yunos.tvpubadmin.content.app.queryapp

查询应用信息
*/
type YunosTvpubadminContentAppQueryappAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAppQueryappResponse
}

type YunosTvpubadminContentAppQueryappResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_queryapp_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Result<AppInfo>
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
