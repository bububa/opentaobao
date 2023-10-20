package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardVirtualphoneBindAPIResponse 工单维度虚拟中间号绑定 API返回值
// tmall.servicecenter.workcard.virtualphone.bind
//
// 服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
// 叫消费者的手机号虚拟号码给到客服。
type TmallServicecenterWorkcardVirtualphoneBindAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardVirtualphoneBindAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardVirtualphoneBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardVirtualphoneBindAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardVirtualphoneBindAPIResponseModel is 工单维度虚拟中间号绑定 成功返回结果
type TmallServicecenterWorkcardVirtualphoneBindAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_virtualphone_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardVirtualphoneBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardVirtualphoneBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardVirtualphoneBindAPIResponse)
	},
}

// GetTmallServicecenterWorkcardVirtualphoneBindAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardVirtualphoneBindAPIResponse
func GetTmallServicecenterWorkcardVirtualphoneBindAPIResponse() *TmallServicecenterWorkcardVirtualphoneBindAPIResponse {
	return poolTmallServicecenterWorkcardVirtualphoneBindAPIResponse.Get().(*TmallServicecenterWorkcardVirtualphoneBindAPIResponse)
}

// ReleaseTmallServicecenterWorkcardVirtualphoneBindAPIResponse 将 TmallServicecenterWorkcardVirtualphoneBindAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardVirtualphoneBindAPIResponse(v *TmallServicecenterWorkcardVirtualphoneBindAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardVirtualphoneBindAPIResponse.Put(v)
}
