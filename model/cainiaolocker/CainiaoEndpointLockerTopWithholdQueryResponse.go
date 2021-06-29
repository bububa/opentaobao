package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询能否代扣 API返回值 
cainiao.endpoint.locker.top.withhold.query

查询是否有代扣欠款，是否签署代扣协议。
*/
type CainiaoEndpointLockerTopWithholdQueryAPIResponse struct {
    model.CommonResponse
    CainiaoEndpointLockerTopWithholdQueryResponse
}

// 查询能否代扣 成功返回结果
type CainiaoEndpointLockerTopWithholdQueryResponse struct {
    XMLName xml.Name `xml:"cainiao_endpoint_locker_top_withhold_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // response
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
