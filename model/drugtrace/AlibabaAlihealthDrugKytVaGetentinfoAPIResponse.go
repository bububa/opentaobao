package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytVaGetentinfoAPIResponse 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API返回值
// alibaba.alihealth.drug.kyt.va.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
type AlibabaAlihealthDrugKytVaGetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytVaGetentinfoAPIResponseModel
}

// AlibabaAlihealthDrugKytVaGetentinfoAPIResponseModel is 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) 成功返回结果
type AlibabaAlihealthDrugKytVaGetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_va_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytVaGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
