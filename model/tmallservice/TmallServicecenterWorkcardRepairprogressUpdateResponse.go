package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新维修进度 APIResponse
tmall.servicecenter.workcard.repairprogress.update

提供给外部合作服务商的维修进度更改接口
*/
type TmallServicecenterWorkcardRepairprogressUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_repairprogress_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *TmallServicecenterWorkcardRepairprogressUpdateResult `json:"result,omitempty" xml:"