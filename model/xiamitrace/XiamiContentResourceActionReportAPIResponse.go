package xiamitrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentResourceActionReportAPIResponse 曲库开放平台内容行为上报接口 API返回值
// xiami.content.resource.action.report
//
// 合作方对接入的曲库开放内容上报行为日志
type XiamiContentResourceActionReportAPIResponse struct {
	model.CommonResponse
	XiamiContentResourceActionReportAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentResourceActionReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentResourceActionReportAPIResponseModel).Reset()
}

// XiamiContentResourceActionReportAPIResponseModel is 曲库开放平台内容行为上报接口 成功返回结果
type XiamiContentResourceActionReportAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_resource_action_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 上报结果: true(成功), false(失败)
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentResourceActionReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = nil
	m.Result = false
}

var poolXiamiContentResourceActionReportAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentResourceActionReportAPIResponse)
	},
}

// GetXiamiContentResourceActionReportAPIResponse 从 sync.Pool 获取 XiamiContentResourceActionReportAPIResponse
func GetXiamiContentResourceActionReportAPIResponse() *XiamiContentResourceActionReportAPIResponse {
	return poolXiamiContentResourceActionReportAPIResponse.Get().(*XiamiContentResourceActionReportAPIResponse)
}

// ReleaseXiamiContentResourceActionReportAPIResponse 将 XiamiContentResourceActionReportAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentResourceActionReportAPIResponse(v *XiamiContentResourceActionReportAPIResponse) {
	v.Reset()
	poolXiamiContentResourceActionReportAPIResponse.Put(v)
}
