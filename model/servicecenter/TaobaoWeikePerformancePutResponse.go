package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交客服绩效接口 APIResponse
taobao.weike.performance.put

提交客服绩效接口
*/
type TaobaoWeikePerformancePutAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"weike_performance_put_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   bool `json:"result,omitempty" xml:"