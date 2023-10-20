package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExpressorderCreateAPIResponse 物流订单创建API API返回值
// tmall.servicecenter.workcard.expressorder.create
//
// 天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
type TmallServicecenterWorkcardExpressorderCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardExpressorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardExpressorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardExpressorderCreateAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardExpressorderCreateAPIResponseModel is 物流订单创建API 成功返回结果
type TmallServicecenterWorkcardExpressorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_expressorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardExpressorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardExpressorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardExpressorderCreateAPIResponse)
	},
}

// GetTmallServicecenterWorkcardExpressorderCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardExpressorderCreateAPIResponse
func GetTmallServicecenterWorkcardExpressorderCreateAPIResponse() *TmallServicecenterWorkcardExpressorderCreateAPIResponse {
	return poolTmallServicecenterWorkcardExpressorderCreateAPIResponse.Get().(*TmallServicecenterWorkcardExpressorderCreateAPIResponse)
}

// ReleaseTmallServicecenterWorkcardExpressorderCreateAPIResponse 将 TmallServicecenterWorkcardExpressorderCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardExpressorderCreateAPIResponse(v *TmallServicecenterWorkcardExpressorderCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardExpressorderCreateAPIResponse.Put(v)
}
