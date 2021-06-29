package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
应用升级查询 APIResponse
yunos.tvpubadmin.device.appupgradequery

应用升级查询
*/
type YunosTvpubadminDeviceAppupgradequeryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceAppupgradequeryResponse
}

type YunosTvpubadminDeviceAppupgradequeryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_appupgradequery_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 具体的数据结构
    
    ObjectList   *PaginationDO `json:"object_list,omitempty" xml:"object_list,omitempty"`

    
}
