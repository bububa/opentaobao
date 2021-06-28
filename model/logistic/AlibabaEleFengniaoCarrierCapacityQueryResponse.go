package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按照门店查询骑手运力状态查询 APIResponse
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
*/
type AlibabaEleFengniaoCarrierCapacityQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_carrier_capacity_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Results   []Capacities `json:"results,omitempty" xml:"