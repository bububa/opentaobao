package cainiaoecc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常统计信息获取 APIResponse
cainiao.ecc.exceptions.delay.count

菜鸟控制塔包裹滞留异常统计信息获取
*/
type CainiaoEccExceptionsDelayCountAPIResponse struct {
    model.CommonResponse
    CainiaoEccExceptionsDelayCountResponse
}

type CainiaoEccExceptionsDelayCountResponse struct {
    XMLName xml.Name `xml:"cainiao_ecc_exceptions_delay_count_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
