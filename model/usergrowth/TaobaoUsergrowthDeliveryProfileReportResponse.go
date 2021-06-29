package usergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标签上报 APIResponse
taobao.usergrowth.delivery.profile.report

渠道上报标签信息
*/
type TaobaoUsergrowthDeliveryProfileReportAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDeliveryProfileReportResponse
}

type TaobaoUsergrowthDeliveryProfileReportResponse struct {
    XMLName xml.Name `xml:"usergrowth_delivery_profile_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
