package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenWavenumReportAPIResponse 发货单波次通知接口 API返回值
// taobao.qimen.wavenum.report
//
// WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoQimenWavenumReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenWavenumReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenWavenumReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenWavenumReportAPIResponseModel).Reset()
}

// TaobaoQimenWavenumReportAPIResponseModel is 发货单波次通知接口 成功返回结果
type TaobaoQimenWavenumReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_wavenum_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenWavenumReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenWavenumReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenWavenumReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWavenumReportAPIResponse)
	},
}

// GetTaobaoQimenWavenumReportAPIResponse 从 sync.Pool 获取 TaobaoQimenWavenumReportAPIResponse
func GetTaobaoQimenWavenumReportAPIResponse() *TaobaoQimenWavenumReportAPIResponse {
	return poolTaobaoQimenWavenumReportAPIResponse.Get().(*TaobaoQimenWavenumReportAPIResponse)
}

// ReleaseTaobaoQimenWavenumReportAPIResponse 将 TaobaoQimenWavenumReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenWavenumReportAPIResponse(v *TaobaoQimenWavenumReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenWavenumReportAPIResponse.Put(v)
}
