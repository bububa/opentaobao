package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSearchbillDetailAPIResponse 查询单据详情 API返回值
// alibaba.alihealth.drug.kyt.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytSearchbillDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel
}

// AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel is 查询单据详情 成功返回结果
type AlibabaAlihealthDrugKytSearchbillDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_searchbill_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSearchbillDetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
