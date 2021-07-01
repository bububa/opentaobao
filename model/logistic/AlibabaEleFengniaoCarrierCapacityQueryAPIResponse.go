package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按照门店查询骑手运力状态查询 API返回值 
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
*/
type AlibabaEleFengniaoCarrierCapacityQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel
}

// 按照门店查询骑手运力状态查询 成功返回结果
type AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_carrier_capacity_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Results   []Capacities `json:"results,omitempty" xml:"results>capacities,omitempty"`
}
