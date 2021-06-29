package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发布应用升级 APIResponse
yunos.osupdate.appversion.publish

发布应用升级任务
*/
type YunosOsupdateAppversionPublishAPIResponse struct {
    model.CommonResponse
    YunosOsupdateAppversionPublishResponse
}

type YunosOsupdateAppversionPublishResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_appversion_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *YunosOsupdateAppversionPublishResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
