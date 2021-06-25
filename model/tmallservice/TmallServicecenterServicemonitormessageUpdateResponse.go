package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商更新预警消息状态 APIResponse
tmall.servicecenter.servicemonitormessage.update

服务商收到预警后，需要进行回复已读状态，并可填写备注
*/
type TmallServicecenterServicemonitormessageUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServicemonitormessageUpdateResponse `json:"tmall_servicecenter_servicemonitormessage_update_response,omitempty"`
}

type TmallServicecenterServicemonitormessageUpdateResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
