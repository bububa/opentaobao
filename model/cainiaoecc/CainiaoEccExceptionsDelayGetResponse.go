package cainiaoecc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常信息获取 APIResponse
cainiao.ecc.exceptions.delay.get

菜鸟控制塔包裹滞留异常信息获取
*/
type CainiaoEccExceptionsDelayGetAPIResponse struct {
    model.CommonResponse
    CainiaoEccExceptionsDelayGetResponse
}

type CainiaoEccExceptionsDelayGetResponse struct {
    XMLName xml.Name `xml:"cainiao_ecc_exceptions_delay_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
