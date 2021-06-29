package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
机型检索 APIResponse
yunos.osupdate.model.search

机型检索
*/
type YunosOsupdateModelSearchAPIResponse struct {
    model.CommonResponse
    YunosOsupdateModelSearchResponse
}

type YunosOsupdateModelSearchResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_model_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 机型列表
    
    ModelList   []DeviceEntryDTO `json:"model_list,omitempty" xml:"model_list>device_entry_dto,omitempty"`
    
    
}
