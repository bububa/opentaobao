package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
机型检索 API返回值 
yunos.osupdate.model.search

机型检索
*/
type YunosOsupdateModelSearchAPIResponse struct {
    model.CommonResponse
    YunosOsupdateModelSearchAPIResponseModel
}

// 机型检索 成功返回结果
type YunosOsupdateModelSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_osupdate_model_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 机型列表
    ModelList   []DeviceEntryDto `json:"model_list,omitempty" xml:"model_list>device_entry_dto,omitempty"`
}
