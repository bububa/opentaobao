package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级发布 APIResponse
yunos.osupdate.osfota.publish

发布osupdate系统升级任务
*/
type YunosOsupdateOsfotaPublishAPIResponse struct {
    model.CommonResponse
    YunosOsupdateOsfotaPublishResponse
}

type YunosOsupdateOsfotaPublishResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_osfota_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *YunosOsupdateOsfotaPublishResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
