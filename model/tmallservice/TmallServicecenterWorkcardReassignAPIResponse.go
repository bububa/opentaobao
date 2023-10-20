package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardreassignAPIResponse 工单改派门店 API返回值
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
type TmallservicecenterworkcardreassignAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardreassignAPIResponseModel
}

// TmallservicecenterworkcardreassignAPIResponseModel is 工单改派门店 成功返回结果
type TmallservicecenterworkcardreassignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_reassign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *TmallservicecenterworkcardreassignResult `json:"result,omitempty" xml:"result,omitempty"`
}
