package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新物流进度 APIResponse
tmall.servicecenter.workcard.updatelogistics

提供给外部合作服务商的物流进度更改接口
*/
type TmallServicecenterWorkcardUpdatelogisticsAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardUpdatelogisticsResponse `json:"tmall_servicecenter_workcard_updatelogistics_response,omitempty"` 
    TmallServicecenterWorkcardUpdatelogisticsResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardUpdatelogisticsResponse struct {

    // 返回信息
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardUpdatelogisticsResponse struct {

    // 返回信息
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
