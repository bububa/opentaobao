package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardrepairprogressupdateAPIResponse 更新维修进度 API返回值
// tmall.servicecenter.workcard.repairprogress.update
//
// 提供给外部合作服务商的维修进度更改接口
type TmallservicecenterworkcardrepairprogressupdateAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardrepairprogressupdateAPIResponseModel
}

// TmallservicecenterworkcardrepairprogressupdateAPIResponseModel is 更新维修进度 成功返回结果
type TmallservicecenterworkcardrepairprogressupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_repairprogress_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallservicecenterworkcardrepairprogressupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
