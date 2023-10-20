package foodscan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafootscanminireportfragmentfirstAPIResponse 第一只脚生成报告接口 API返回值
// alibaba.footscan.mini.report.fragment.first
//
// 第一只脚生成报告接口
type AlibabafootscanminireportfragmentfirstAPIResponse struct {
	model.CommonResponse
	AlibabafootscanminireportfragmentfirstAPIResponseModel
}

// AlibabafootscanminireportfragmentfirstAPIResponseModel is 第一只脚生成报告接口 成功返回结果
type AlibabafootscanminireportfragmentfirstAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_footscan_mini_report_fragment_first_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AlibabafootscanminireportfragmentfirstMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
