package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardcallrecordAPIResponse 回访记录回传API API返回值
// tmall.servicecenter.workcard.call.record
//
// 客满回访信息登记
type TmallservicecenterworkcardcallrecordAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardcallrecordAPIResponseModel
}

// TmallservicecenterworkcardcallrecordAPIResponseModel is 回访记录回传API 成功返回结果
type TmallservicecenterworkcardcallrecordAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_call_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
