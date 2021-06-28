package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新物料制作状态 APIResponse
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态
*/
type AlibabaCuntaoInteractRequisitionUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaCuntaoInteractRequisitionUpdateResponse `json:"alibaba_cuntao_interact_requisition_update_response,omitempty"` 
    AlibabaCuntaoInteractRequisitionUpdateResponse
}

/* model for simplify = false
type AlibabaCuntaoInteractRequisitionUpdateResponse struct {

    // result
    
    Result  *struct {
        AlibabaCuntaoInteractRequisitionUpdateResult  *AlibabaCuntaoInteractRequisitionUpdateResult `json:"alibaba_cuntao_interact_requisition_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaCuntaoInteractRequisitionUpdateResponse struct {

    // result
    Result   *AlibabaCuntaoInteractRequisitionUpdateResult `json:"result,omitempty"`

}
