package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 APIResponse
yunos.tvpubadmin.device.appupgradedetail

获取应用升级详情
*/
type YunosTvpubadminDeviceAppupgradedetailAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceAppupgradedetailResponse
}

type YunosTvpubadminDeviceAppupgradedetailResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_appupgradedetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 应用列表
    
    Object   *AppVersionAuditDo `json:"object,omitempty" xml:"object,omitempty"`

    
}
