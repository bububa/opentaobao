package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDestbillListAPIResponse 直调单据查询 API返回值
// alibaba.alihealth.drug.kyt.destbill.list
//
// 为药企提供直调单据查询功能
type AlibabaAlihealthDrugKytDestbillListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDestbillListAPIResponseModel
}

// AlibabaAlihealthDrugKytDestbillListAPIResponseModel is 直调单据查询 成功返回结果
type AlibabaAlihealthDrugKytDestbillListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_destbill_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回result
	Result *AlibabaAlihealthDrugKytDestbillListResult `json:"result,omitempty" xml:"result,omitempty"`
}
