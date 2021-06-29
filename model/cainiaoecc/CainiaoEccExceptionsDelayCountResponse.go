package cainiaoecc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常统计信息获取 API返回值 
cainiao.ecc.exceptions.delay.count

菜鸟控制塔包裹滞留异常统计信息获取
*/
type CainiaoEccExceptionsDelayCountAPIResponse struct {
    model.CommonResponse
    CainiaoEccExceptionsDelayCountResponse
}

// 菜鸟控制塔包裹滞留异常统计信息获取 成功返回结果
type CainiaoEccExceptionsDelayCountResponse struct {
    XMLName xml.Name `xml:"cainiao_ecc_exceptions_delay_count_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
