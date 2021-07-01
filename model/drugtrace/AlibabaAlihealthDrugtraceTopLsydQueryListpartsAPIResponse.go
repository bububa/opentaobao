package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse
往来单位查询 API返回值
alibaba.alihealth.drugtrace.top.lsyd.query.listparts

查询往来单位列表 */
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel
}

// AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel is 往来单位查询 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
