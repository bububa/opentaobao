package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyputpackagebindAPIResponse 码拼箱建立父子关系接口 API返回值
// alibaba.alihealth.drug.kyt.scqy.putpackagebind
//
// 码拼箱建立父子关系接口
type AlibabaalihealthdrugkytscqyputpackagebindAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytscqyputpackagebindAPIResponseModel
}

// AlibabaalihealthdrugkytscqyputpackagebindAPIResponseModel is 码拼箱建立父子关系接口 成功返回结果
type AlibabaalihealthdrugkytscqyputpackagebindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_putpackagebind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaalihealthdrugkytscqyputpackagebindResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
