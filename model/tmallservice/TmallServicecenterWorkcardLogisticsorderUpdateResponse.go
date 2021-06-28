package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流工单履约状态更新 APIResponse
tmall.servicecenter.workcard.logisticsorder.update

天猫寄送类服务对接外部物流服务商回传物流状态信息
*/
type TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardLogisticsorderUpdateResponse
}

type TmallServicecenterWorkcardLogisticsorderUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsorder_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
