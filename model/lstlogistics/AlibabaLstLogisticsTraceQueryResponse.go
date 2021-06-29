package lstlogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-查询运单物流追踪信息 API返回值 
alibaba.lst.logistics.trace.query

查询LP单物流追踪信息
*/
type AlibabaLstLogisticsTraceQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstLogisticsTraceQueryResponse
}

// 供应商-异云-查询运单物流追踪信息 成功返回结果
type AlibabaLstLogisticsTraceQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_logistics_trace_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaLstLogisticsTraceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
