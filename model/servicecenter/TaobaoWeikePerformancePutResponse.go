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
    Response *TaobaoWeikePerformancePutResponse `json:"taobao_weike_performance_put_response,omitempty"`
}

type TaobaoWeikePerformancePutResponse struct {

    // 返回结果
    Result   bool `json:"result,omitempty"`

}
