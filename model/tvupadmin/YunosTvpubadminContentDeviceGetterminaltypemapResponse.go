package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取终端类型表 APIResponse
yunos.tvpubadmin.content.device.getterminaltypemap

获取终端类型表
*/
type YunosTvpubadminContentDeviceGetterminaltypemapAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentDeviceGetterminaltypemapResponse
}

type YunosTvpubadminContentDeviceGetterminaltypemapResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_device_getterminaltypemap_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // map
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
