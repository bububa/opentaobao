package tmallservice

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallServicecenterWorkcardConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardConfirmAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardConfirmAPIResponseModel is 服务商确认服务完成 成功返回结果
type TmallServicecenterWorkcardConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardConfirmAPIResponse)
	},
}

// GetTmallServicecenterWorkcardConfirmAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardConfirmAPIResponse
func GetTmallServicecenterWorkcardConfirmAPIResponse() *TmallServicecenterWorkcardConfirmAPIResponse {
	return poolTmallServicecenterWorkcardConfirmAPIResponse.Get().(*TmallServicecenterWorkcardConfirmAPIResponse)
}

// ReleaseTmallServicecenterWorkcardConfirmAPIResponse 将 TmallServicecenterWorkcardConfirmAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardConfirmAPIResponse(v *TmallServicecenterWorkcardConfirmAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardConfirmAPIResponse.Put(v)
}
