package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
节目审核获取节目列表 APIResponse
yunos.tvpubadmin.content.show.getlist

节目审核获取节目列表
*/
type YunosTvpubadminContentShowGetlistAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentShowGetlistResponse
}

type YunosTvpubadminContentShowGetlistResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
