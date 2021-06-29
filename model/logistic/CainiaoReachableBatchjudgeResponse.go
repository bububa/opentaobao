package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
是否派送可达判定批量查询接口 API返回值 
cainiao.reachable.batchjudge

提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
*/
type CainiaoReachableBatchjudgeAPIResponse struct {
    model.CommonResponse
    CainiaoReachableBatchjudgeResponse
}

// 是否派送可达判定批量查询接口 成功返回结果
type CainiaoReachableBatchjudgeResponse struct {
    XMLName xml.Name `xml:"cainiao_reachable_batchjudge_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
