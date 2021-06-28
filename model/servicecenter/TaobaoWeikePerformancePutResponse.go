package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提交客服绩效接口 APIResponse
taobao.weike.performance.put

提交客服绩效接口
*/
type TaobaoWeikePerformancePutAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWeikePerformancePutResponse `json:"weike_performance_put_response,omitempty"` 
    TaobaoWeikePerformancePutResponse
}

/* model for simplify = false
type TaobaoWeikePerformancePutResponse struct {

    // 返回结果
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoWeikePerformancePutResponse struct {

    // 返回结果
    Result   bool `json:"result,omitempty"`

}
