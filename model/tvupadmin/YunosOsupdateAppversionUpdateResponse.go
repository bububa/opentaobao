package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
应用升级任务更新 APIResponse
yunos.osupdate.appversion.update

应用升级任务更新
*/
type YunosOsupdateAppversionUpdateAPIResponse struct {
    model.CommonResponse
    YunosOsupdateAppversionUpdateResponse
}

type YunosOsupdateAppversionUpdateResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_appversion_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
