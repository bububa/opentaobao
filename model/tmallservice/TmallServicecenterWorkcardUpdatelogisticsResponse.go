package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新物流进度 APIResponse
tmall.servicecenter.workcard.updatelogistics

提供给外部合作服务商的物流进度更改接口
*/
type TmallServicecenterWorkcardUpdatelogisticsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_updatelogistics_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回信息
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"