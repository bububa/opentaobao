package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmidslistAPIResponse 同情用药根据申请单列表查询申请单 API返回值
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
type AlibabaalihealthpwgmidslistAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthpwgmidslistAPIResponseModel
}

// AlibabaalihealthpwgmidslistAPIResponseModel is 同情用药根据申请单列表查询申请单 成功返回结果
type AlibabaalihealthpwgmidslistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_ids_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
