package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务列表 APIResponse
yunos.tvpubadmin.diccontroltask.query

牌照方对终端设备的停开服管理
*/
type YunosTvpubadminDiccontroltaskQueryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDiccontroltaskQueryResponse
}

type YunosTvpubadminDiccontroltaskQueryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *PaginationDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
