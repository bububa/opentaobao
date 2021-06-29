package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
1-IBP同步物料接口 APIResponse
alibaba.tmallgenie.scp.plan.materiel.get

IBP同步物料接口
*/
type AlibabaTmallgenieScpPlanMaterielGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanMaterielGetResponse
}

type AlibabaTmallgenieScpPlanMaterielGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_materiel_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
