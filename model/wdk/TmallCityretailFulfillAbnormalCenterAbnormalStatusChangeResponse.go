package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
同城零售履约异常中心异常单处理结果回调接口 APIResponse
tmall.cityretail.fulfill.abnormal.center.abnormal.status.change

同城零售履约异常中心异常单处理结果回调接口
*/
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse struct {
    model.CommonResponse
    Response *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeResponse `json:"tmall_cityretail_fulfill_abnormal_center_abnormal_status_change_response,omitempty"`
}

type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeResponse struct {

    // 响应参数
    FulfillSingleResult   *FulfillSingleResult `json:"fulfill_single_result,omitempty"`

}
