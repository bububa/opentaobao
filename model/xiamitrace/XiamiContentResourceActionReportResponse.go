package xiamitrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
曲库开放平台内容行为上报接口 APIResponse
xiami.content.resource.action.report

合作方对接入的曲库开放内容上报行为日志
*/
type XiamiContentResourceActionReportAPIResponse struct {
    model.CommonResponse
    XiamiContentResourceActionReportResponse
}

type XiamiContentResourceActionReportResponse struct {
    XMLName xml.Name `xml:"xiami_content_resource_action_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 上报结果: true(成功), false(失败)
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 返回结果码
    
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
