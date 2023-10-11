package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse 物流代货主查找往来单位接口 API返回值
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel
}

// AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel is 物流代货主查找往来单位接口 成功返回结果
type AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_listparts_byagent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesListpartsByagentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
