package healthnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新零售物流详情接口 API返回值 
alibaba.health.nr.logistics.query

对阿里健康o2o对接的商户提供查询物流单详情的能力
*/
type AlibabaHealthNrLogisticsQueryAPIResponse struct {
    model.CommonResponse
    AlibabaHealthNrLogisticsQueryResponse
}

// 阿里健康新零售物流详情接口 成功返回结果
type AlibabaHealthNrLogisticsQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_health_nr_logistics_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    ResponseResult   *ResponseResult `json:"response_result,omitempty" xml:"response_result,omitempty"`
}
