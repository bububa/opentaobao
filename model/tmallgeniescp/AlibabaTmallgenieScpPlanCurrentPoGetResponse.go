package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
11-同步本周的po单（从W-1周到W+4周） APIResponse
alibaba.tmallgenie.scp.plan.current.po.get

11-同步本周的po单（从W-1周到W+4周）
*/
type AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanCurrentPoGetResponse
}

type AlibabaTmallgenieScpPlanCurrentPoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_current_po_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
