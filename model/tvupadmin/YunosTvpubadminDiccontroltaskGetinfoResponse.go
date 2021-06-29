package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服任务详情 APIResponse
yunos.tvpubadmin.diccontroltask.getinfo

获取停开服任务详情
*/
type YunosTvpubadminDiccontroltaskGetinfoAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDiccontroltaskGetinfoResponse
}

type YunosTvpubadminDiccontroltaskGetinfoResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_getinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *DicControlTaskDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
