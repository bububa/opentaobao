package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取系统升级详情 APIResponse
yunos.tvpubadmin.device.osupgradedetail

获取系统升级详情
*/
type YunosTvpubadminDeviceOsupgradedetailAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceOsupgradedetailResponse
}

type YunosTvpubadminDeviceOsupgradedetailResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_osupgradedetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 具体的数据对象
    
    Object   *OsVersionAuditDo `json:"object,omitempty" xml:"object,omitempty"`

    
}
