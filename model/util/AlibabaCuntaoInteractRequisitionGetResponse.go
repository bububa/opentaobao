package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商获取物料申请单列表 APIResponse
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表
*/
type AlibabaCuntaoInteractRequisitionGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaCuntaoInteractRequisitionGetResponse `json:"alibaba_cuntao_interact_requisition_get_response,omitempty"` 
    AlibabaCuntaoInteractRequisitionGetResponse
}

/* model for simplify = false
type AlibabaCuntaoInteractRequisitionGetResponse struct {

    // result
    
    Result  *struct {
        AlibabaCuntaoInteractRequisitionGetResult  *AlibabaCuntaoInteractRequisitionGetResult `json:"alibaba_cuntao_interact_requisition_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaCuntaoInteractRequisitionGetResponse struct {

    // result
    Result   *AlibabaCuntaoInteractRequisitionGetResult `json:"result,omitempty"`

}
