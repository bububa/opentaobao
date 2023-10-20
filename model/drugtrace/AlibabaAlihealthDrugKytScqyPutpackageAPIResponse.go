package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyputpackageAPIResponse 码拼箱 API返回值
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
type AlibabaalihealthdrugkytscqyputpackageAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytscqyputpackageAPIResponseModel
}

// AlibabaalihealthdrugkytscqyputpackageAPIResponseModel is 码拼箱 成功返回结果
type AlibabaalihealthdrugkytscqyputpackageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_putpackage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaalihealthdrugkytscqyputpackageResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
