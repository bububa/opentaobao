package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建应用升级任务 APIResponse
yunos.osupdate.appversion.create

创建应用升级任务
*/
type YunosOsupdateAppversionCreateAPIResponse struct {
    model.CommonResponse
    YunosOsupdateAppversionCreateResponse
}

type YunosOsupdateAppversionCreateResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_appversion_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
