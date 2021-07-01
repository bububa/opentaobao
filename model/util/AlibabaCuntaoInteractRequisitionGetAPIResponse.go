package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCuntaoInteractRequisitionGetAPIResponse
供应商获取物料申请单列表 API返回值
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表 */
type AlibabaCuntaoInteractRequisitionGetAPIResponse struct {
	model.CommonResponse
	AlibabaCuntaoInteractRequisitionGetAPIResponseModel
}

// AlibabaCuntaoInteractRequisitionGetAPIResponseModel is 供应商获取物料申请单列表 成功返回结果
type AlibabaCuntaoInteractRequisitionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cuntao_interact_requisition_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaCuntaoInteractRequisitionGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
