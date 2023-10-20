package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljgdrugtableAPIResponse 查询药品目录信息 API返回值
// alibaba.alihealth.drugtrace.top.yljg.drugtable
//
// 查询药品目录信息
type AlibabaalihealthdrugtracetopyljgdrugtableAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetopyljgdrugtableAPIResponseModel
}

// AlibabaalihealthdrugtracetopyljgdrugtableAPIResponseModel is 查询药品目录信息 成功返回结果
type AlibabaalihealthdrugtracetopyljgdrugtableAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_drugtable_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugtracetopyljgdrugtableResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
