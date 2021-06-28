package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按照门店查询骑手运力状态查询 APIResponse
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
*/
type AlibabaEleFengniaoCarrierCapacityQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaEleFengniaoCarrierCapacityQueryResponse `json:"alibaba_ele_fengniao_carrier_capacity_query_response,omitempty"` 
    AlibabaEleFengniaoCarrierCapacityQueryResponse
}

/* model for simplify = false
type AlibabaEleFengniaoCarrierCapacityQueryResponse struct {

    // 系统自动生成
    
    Results  struct {
        Capacities  []Capacities `json:"capacities,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type AlibabaEleFengniaoCarrierCapacityQueryResponse struct {

    // 系统自动生成
    Results   []Capacities `json:"results,omitempty"`

}
