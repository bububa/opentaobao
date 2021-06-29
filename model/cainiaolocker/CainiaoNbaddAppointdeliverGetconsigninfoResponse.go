package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支持定时派送服务发货信息 APIResponse
cainiao.nbadd.appointdeliver.getconsigninfo

获取支持定时派送服务发货信息
*/
type CainiaoNbaddAppointdeliverGetconsigninfoAPIResponse struct {
    model.CommonResponse
    CainiaoNbaddAppointdeliverGetconsigninfoResponse
}

type CainiaoNbaddAppointdeliverGetconsigninfoResponse struct {
    XMLName xml.Name `xml:"cainiao_nbadd_appointdeliver_getconsigninfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用是否成功，正常情况下都会成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误编码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误描述
    
    ResultDesc   string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`

    
    // 发货信息
    
    Result   *ConsignSupportInfoDTO `json:"result,omitempty" xml:"result,omitempty"`

    
}
