package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardConfirmAPIResponse 服务商确认服务完成 API返回值
// tmall.servicecenter.workcard.confirm
//
// 提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
type TmallServicecenterWorkcardConfirmAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardConfirmAPIResponseModel
}

// TmallServicecenterWorkcardConfirmAPIResponseModel is 服务商确认服务完成 成功返回结果
type TmallServicecenterWorkcardConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
