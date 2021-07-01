package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCuntaoInteractRequisitionUpdateAPIResponse
更新物料制作状态 API返回值
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态 */
type AlibabaCuntaoInteractRequisitionUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaCuntaoInteractRequisitionUpdateAPIResponseModel
}

// AlibabaCuntaoInteractRequisitionUpdateAPIResponseModel is 更新物料制作状态 成功返回结果
type AlibabaCuntaoInteractRequisitionUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cuntao_interact_requisition_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaCuntaoInteractRequisitionUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
