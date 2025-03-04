package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse UD效果外投投放监测链接生成 API返回值
// alibaba.taobao.ud.smart.monitor.url.create
//
// UD效果外投投放监测链接生成
type AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponseModel).Reset()
}

// AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponseModel is UD效果外投投放监测链接生成 成功返回结果
type AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_taobao_ud_smart_monitor_url_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监测链接信息
	Result *MonitorUrlTopDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse)
	},
}

// GetAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse 从 sync.Pool 获取 AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse
func GetAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse() *AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse {
	return poolAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse.Get().(*AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse)
}

// ReleaseAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse 将 AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse(v *AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse) {
	v.Reset()
	poolAlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse.Put(v)
}
