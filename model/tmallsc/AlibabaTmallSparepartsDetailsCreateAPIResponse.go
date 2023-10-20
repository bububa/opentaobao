package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallsparepartsdetailscreateAPIResponse 天猫蚁巢同步工单申请备件明细 API返回值
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
type AlibabatmallsparepartsdetailscreateAPIResponse struct {
	model.CommonResponse
	AlibabatmallsparepartsdetailscreateAPIResponseModel
}

// AlibabatmallsparepartsdetailscreateAPIResponseModel is 天猫蚁巢同步工单申请备件明细 成功返回结果
type AlibabatmallsparepartsdetailscreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmall_spareparts_details_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	DisplayMessage string `json:"display_message,omitempty" xml:"display_message,omitempty"`
	// app名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误参数
	ErrorParams string `json:"error_params,omitempty" xml:"error_params,omitempty"`
	// 返回数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
