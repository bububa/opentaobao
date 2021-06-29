package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据活动id查询活动相关快闪店及设备信息 APIResponse
tmall.popupstore.activity.device.query

查询某一活动的deviceCode的部署情况
*/
type TmallPopupstoreActivityDeviceQueryAPIResponse struct {
    model.CommonResponse
    TmallPopupstoreActivityDeviceQueryResponse
}

type TmallPopupstoreActivityDeviceQueryResponse struct {
    XMLName xml.Name `xml:"tmall_popupstore_activity_device_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参对象
    
    ResultDto   *TmallPopupstoreActivityDeviceQueryResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`

    
}
