package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货pv查询 API返回值 
alibaba.idle.isv.pv.list

根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
*/
type AlibabaIdleIsvPvListAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvPvListAPIResponseModel
}

// 闲鱼已验货pv查询 成功返回结果
type AlibabaIdleIsvPvListAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_pv_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaIdleIsvPvListResult `json:"result,omitempty" xml:"result,omitempty"`
}
