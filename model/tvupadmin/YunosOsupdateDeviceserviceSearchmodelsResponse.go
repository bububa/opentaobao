package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据关键词检索设备型号 APIResponse
yunos.osupdate.deviceservice.searchmodels

根据关键词检索设备型号
*/
type YunosOsupdateDeviceserviceSearchmodelsAPIResponse struct {
    model.CommonResponse
    YunosOsupdateDeviceserviceSearchmodelsResponse
}

type YunosOsupdateDeviceserviceSearchmodelsResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_deviceservice_searchmodels_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    ModelList   []DeviceEntryDTO `json:"model_list,omitempty" xml:"model_list>device_entry_dto,omitempty"`
    
    
}
