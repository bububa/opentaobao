package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse 获取服务截止日期 API返回值
// alibaba.alihealth.drugtrace.top.yljg.service.getenddate
//
// 获取企业服务截止时间
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel
}

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel is 获取服务截止日期 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_service_getenddate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
