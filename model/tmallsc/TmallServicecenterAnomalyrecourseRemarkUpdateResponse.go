package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台一键求助单服务商备注更新接口 APIResponse
tmall.servicecenter.anomalyrecourse.remark.update

一键求助服务商可以回传备注
*/
type TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterAnomalyrecourseRemarkUpdateResponse `json:"tmall_servicecenter_anomalyrecourse_remark_update_response,omitempty"`
}

type TmallServicecenterAnomalyrecourseRemarkUpdateResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
