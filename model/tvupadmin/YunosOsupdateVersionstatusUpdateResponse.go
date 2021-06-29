package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用升级状态 API返回值 
yunos.osupdate.versionstatus.update

更新应用升级状态
*/
type YunosOsupdateVersionstatusUpdateAPIResponse struct {
    model.CommonResponse
    YunosOsupdateVersionstatusUpdateResponse
}

// 更新应用升级状态 成功返回结果
type YunosOsupdateVersionstatusUpdateResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_versionstatus_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
