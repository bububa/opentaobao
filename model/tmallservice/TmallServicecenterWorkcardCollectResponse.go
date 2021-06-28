package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单揽件 APIResponse
tmall.servicecenter.workcard.collect

服务工单揽件接口
*/
type TmallServicecenterWorkcardCollectAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_collect_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"