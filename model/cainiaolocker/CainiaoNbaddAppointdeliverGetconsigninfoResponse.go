package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支持定时派送服务发货信息 API返回值 
cainiao.nbadd.appointdeliver.getconsigninfo

获取支持定时派送服务发货信息
*/
type CainiaoNbaddAppointdeliverGetconsigninfoAPIResponse struct {
    model.CommonResponse
    CainiaoNbaddAppointdeliverGetconsigninfoResponse
}

// 获取支持定时派送服务发货信息 成功返回结果
type CainiaoNbaddAppointdeliverGetconsigninfoResponse struct {
    XMLName xml.Name `xml:"cainiao_nbadd_appointdeliver_getconsigninfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用是否成功，正常情况下都会成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 错误编码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 错误描述
    ResultDesc   string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
    // 发货信息
    Result   *ConsignSupportInfoDto `json:"result,omitempty" xml:"result,omitempty"`
}
