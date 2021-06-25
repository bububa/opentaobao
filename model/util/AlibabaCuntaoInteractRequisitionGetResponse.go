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
    Response *AlibabaCuntaoInteractRequisitionGetResponse `json:"alibaba_cuntao_interact_requisition_get_response,omitempty"`
}

type AlibabaCuntaoInteractRequisitionGetResponse struct {

    // result
    Result   *AlibabaCuntaoInteractRequisitionGetResult `json:"result,omitempty"`

}
