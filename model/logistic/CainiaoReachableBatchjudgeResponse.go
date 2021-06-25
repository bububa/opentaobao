package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
是否派送可达判定批量查询接口 APIResponse
cainiao.reachable.batchjudge

提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
*/
type CainiaoReachableBatchjudgeAPIResponse struct {
    model.CommonResponse
    Response *CainiaoReachableBatchjudgeResponse `json:"cainiao_reachable_batchjudge_response,omitempty"`
}

type CainiaoReachableBatchjudgeResponse struct {

    // 结果
    Result   *BaseResultDto `json:"result,omitempty"`

}