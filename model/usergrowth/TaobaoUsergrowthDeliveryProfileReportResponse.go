package usergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标签上报 API返回值 
taobao.usergrowth.delivery.profile.report

渠道上报标签信息
*/
type TaobaoUsergrowthDeliveryProfileReportAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDeliveryProfileReportResponse
}

// 标签上报 成功返回结果
type TaobaoUsergrowthDeliveryProfileReportResponse struct {
    XMLName xml.Name `xml:"usergrowth_delivery_profile_report_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
