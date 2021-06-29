package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增停开服任务 APIResponse
yunos.tvpubadmin.diccontroltask.add

新增停开服任务
*/
type YunosTvpubadminDiccontroltaskAddAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDiccontroltaskAddResponse
}

type YunosTvpubadminDiccontroltaskAddResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
