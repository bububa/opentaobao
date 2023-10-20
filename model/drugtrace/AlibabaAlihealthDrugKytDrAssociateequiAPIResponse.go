package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrassociateequiAPIResponse 疫苗单据与设备绑定 API返回值
// alibaba.alihealth.drug.kyt.dr.associateequi
//
// 疫苗单据与设备绑定
type AlibabaalihealthdrugkytdrassociateequiAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytdrassociateequiAPIResponseModel
}

// AlibabaalihealthdrugkytdrassociateequiAPIResponseModel is 疫苗单据与设备绑定 成功返回结果
type AlibabaalihealthdrugkytdrassociateequiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_associateequi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihealthdrugkytdrassociateequiResult `json:"result,omitempty" xml:"result,omitempty"`
}
