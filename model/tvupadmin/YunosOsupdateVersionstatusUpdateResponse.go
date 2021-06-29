package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用升级状态 APIResponse
yunos.osupdate.versionstatus.update

更新应用升级状态
*/
type YunosOsupdateVersionstatusUpdateAPIResponse struct {
    model.CommonResponse
    YunosOsupdateVersionstatusUpdateResponse
}

type YunosOsupdateVersionstatusUpdateResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_versionstatus_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
