package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponse 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponseModel
}

// AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponseModel is 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_upbilldetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
