package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
云POS查看专柜属性 APIResponse
alibaba.mos.falcon.pos.counter.query

银泰商业获取专柜是否支持小数等属性查看
*/
type AlibabaMosFalconPosCounterQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosFalconPosCounterQueryResponse `json:"alibaba_mos_falcon_pos_counter_query_response,omitempty"`
}

type AlibabaMosFalconPosCounterQueryResponse struct {

    // 返回结果
    Result   *AlibabaMosFalconPosCounterQueryResultDo `json:"result,omitempty"`

}
