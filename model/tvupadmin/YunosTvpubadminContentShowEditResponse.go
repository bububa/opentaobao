package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
媒资节目信息修改 APIResponse
yunos.tvpubadmin.content.show.edit

供迎客松修改媒资节目信息
*/
type YunosTvpubadminContentShowEditAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentShowEditResponse
}

type YunosTvpubadminContentShowEditResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
