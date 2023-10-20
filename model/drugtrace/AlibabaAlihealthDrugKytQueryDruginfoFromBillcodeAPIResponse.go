package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponse 根据单据编号查询单据明细 API返回值
// alibaba.alihealth.drug.kyt.query.druginfo.from.billcode
//
// 根据单据编号查询单据明细
type AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponseModel
}

// AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponseModel is 根据单据编号查询单据明细 成功返回结果
type AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_druginfo_from_billcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytquerydruginfofrombillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
