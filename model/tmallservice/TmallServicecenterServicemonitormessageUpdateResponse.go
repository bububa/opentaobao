package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商更新预警消息状态 APIResponse
tmall.servicecenter.servicemonitormessage.update

服务商收到预警后，需要进行回复已读状态，并可填写备注
*/
type TmallServicecenterServicemonitormessageUpdateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicemonitormessageUpdateResponse
}

type TmallServicecenterServicemonitormessageUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicemonitormessage_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
