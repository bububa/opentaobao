package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流工单履约状态更新 APIResponse
tmall.servicecenter.workcard.logisticsorder.update

天猫寄送类服务对接外部物流服务商回传物流状态信息
*/
type TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardLogisticsorderUpdateResponse `json:"tmall_servicecenter_workcard_logisticsorder_update_response,omitempty"`
}

type TmallServicecenterWorkcardLogisticsorderUpdateResponse struct {

    // 系统自动生成
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
