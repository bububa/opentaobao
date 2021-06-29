package cainiaohandover

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取交接单小包信息 APIResponse
cainiao.global.handover.parcel.query

提供给ISV通过该接口查询小包信息
*/
type CainiaoGlobalHandoverParcelQueryAPIResponse struct {
    model.CommonResponse
    CainiaoGlobalHandoverParcelQueryResponse
}

type CainiaoGlobalHandoverParcelQueryResponse struct {
    XMLName xml.Name `xml:"cainiao_global_handover_parcel_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果
    
    Result   *HsfResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
