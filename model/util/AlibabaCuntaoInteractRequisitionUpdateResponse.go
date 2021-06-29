package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新物料制作状态 APIResponse
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态
*/
type AlibabaCuntaoInteractRequisitionUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaCuntaoInteractRequisitionUpdateResponse
}

type AlibabaCuntaoInteractRequisitionUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_cuntao_interact_requisition_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaCuntaoInteractRequisitionUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
