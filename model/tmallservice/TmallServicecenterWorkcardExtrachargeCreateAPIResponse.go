package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExtrachargeCreateAPIResponse 创建工单额外收费项 API返回值
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
type TmallServicecenterWorkcardExtrachargeCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardExtrachargeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardExtrachargeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardExtrachargeCreateAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardExtrachargeCreateAPIResponseModel is 创建工单额外收费项 成功返回结果
type TmallServicecenterWorkcardExtrachargeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_extracharge_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardExtrachargeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardExtrachargeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardExtrachargeCreateAPIResponse)
	},
}

// GetTmallServicecenterWorkcardExtrachargeCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardExtrachargeCreateAPIResponse
func GetTmallServicecenterWorkcardExtrachargeCreateAPIResponse() *TmallServicecenterWorkcardExtrachargeCreateAPIResponse {
	return poolTmallServicecenterWorkcardExtrachargeCreateAPIResponse.Get().(*TmallServicecenterWorkcardExtrachargeCreateAPIResponse)
}

// ReleaseTmallServicecenterWorkcardExtrachargeCreateAPIResponse 将 TmallServicecenterWorkcardExtrachargeCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardExtrachargeCreateAPIResponse(v *TmallServicecenterWorkcardExtrachargeCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardExtrachargeCreateAPIResponse.Put(v)
}
