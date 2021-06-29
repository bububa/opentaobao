package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松节目设置免审开关 APIResponse
yunos.tvpubadmin.content.show.setexemptaudit

迎客松节目设置免审开关
*/
type YunosTvpubadminContentShowSetexemptauditAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentShowSetexemptauditResponse
}

type YunosTvpubadminContentShowSetexemptauditResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_setexemptaudit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设置免审
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
