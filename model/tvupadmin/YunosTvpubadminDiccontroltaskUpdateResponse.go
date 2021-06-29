package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务状态变更 APIResponse
yunos.tvpubadmin.diccontroltask.update

停开服任务状态变更
*/
type YunosTvpubadminDiccontroltaskUpdateAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDiccontroltaskUpdateResponse
}

type YunosTvpubadminDiccontroltaskUpdateResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
