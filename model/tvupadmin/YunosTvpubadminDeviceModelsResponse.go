package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取品牌下设备列表 APIResponse
yunos.tvpubadmin.device.models

获取品牌下设备列表
*/
type YunosTvpubadminDeviceModelsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceModelsResponse
}

type YunosTvpubadminDeviceModelsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_models_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    ModelList   []string `json:"model_list,omitempty" xml:"model_list>string,omitempty"`
    
    
}
