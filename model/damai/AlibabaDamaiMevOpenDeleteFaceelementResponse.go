package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement APIResponse
alibaba.damai.mev.open.delete.faceelement

deleteFaceElement
*/
type AlibabaDamaiMevOpenDeleteFaceelementAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenDeleteFaceelementResponse
}

type AlibabaDamaiMevOpenDeleteFaceelementResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_delete_faceelement_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenDeleteFaceelementResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
