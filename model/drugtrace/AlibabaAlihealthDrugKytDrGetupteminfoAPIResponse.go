package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrgetupteminfoAPIResponse 获取上游温度信息（疫苗） API返回值
// alibaba.alihealth.drug.kyt.dr.getupteminfo
//
// 根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
type AlibabaalihealthdrugkytdrgetupteminfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytdrgetupteminfoAPIResponseModel
}

// AlibabaalihealthdrugkytdrgetupteminfoAPIResponseModel is 获取上游温度信息（疫苗） 成功返回结果
type AlibabaalihealthdrugkytdrgetupteminfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getupteminfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihealthdrugkytdrgetupteminfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
