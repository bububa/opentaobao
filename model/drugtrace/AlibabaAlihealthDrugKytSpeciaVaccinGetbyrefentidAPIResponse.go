package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponse 根据企业唯一标识查看企业详情 API返回值
// alibaba.alihealth.drug.kyt.specia.vaccin.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponseModel
}

// AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详情 成功返回结果
type AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_specia_vaccin_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytspeciavaccingetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
