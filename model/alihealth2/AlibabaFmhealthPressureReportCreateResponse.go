package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
血压报告接口 APIResponse
alibaba.fmhealth.pressure.report.create

生成用户血压测量报告
*/
type AlibabaFmhealthPressureReportCreateAPIResponse struct {
    model.CommonResponse
    AlibabaFmhealthPressureReportCreateResponse
}

type AlibabaFmhealthPressureReportCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_fmhealth_pressure_report_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgCode
    
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // data
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // success
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`

    
}
